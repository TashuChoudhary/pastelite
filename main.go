package main

import (
	"PasteLite/storage"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/paste", func(c *fiber.Ctx) error {
		type PasteRequest struct {
			Content string `json:"content"`
			Expiry  string `json:"expiry"`
		}
		var req PasteRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid Request")

		}

		duration, err := time.ParseDuration(req.Expiry)
		if err != nil {
			duration = time.Hour * 1

		}

		id := uuid.New().String()
		quote := storage.GetRandomQuote()
		storage.SavePaste(id, req.Content, quote, duration)

		return c.JSON(fiber.Map{
			"id":     id,
			"expiry": duration.String(),
			"url":    fmt.Sprintf("/paste/%s", id),
			"quote":  quote,
		})

	})

	app.Get("/paste/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		paste, found := storage.GetPaste(id)
		if !found {
			return c.Status(fiber.StatusNotFound).SendString("Paste not found")

		}

		if c.Get("Accept") == "text/html" || c.Get("User-Agent") != "" {
			return c.SendFile("./public/paste.html")
		}
		return c.JSON(fiber.Map{
			"id":      id,
			"content": paste.Content,
			"quote":   paste.Quote,
		})

	})
	app.Get("/api/paste/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		paste, found := storage.GetPaste(id)
		if !found {
			return c.Status(404).SendString("Paste not found")
		}
		return c.JSON(fiber.Map{
			"id":      id,
			"content": paste.Content,
			"quote":   paste.Quote,
			"expires": paste.ExpiresAt.Unix(),
		})
	})
	app.Listen(":3000")
}
