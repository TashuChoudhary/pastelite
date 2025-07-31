
# 📝 PasteLite – A Minimal Pastebin Clone

PasteLite is a lightweight web-based application for sharing code or text snippets, inspired by Pastebin. It is built with Go for the backend, a simple HTML/CSS frontend, and supports containerized deployment using Docker and NGINX. The app is deployed using a CI/CD pipeline via GitHub Actions and hosted on Render.

---

## 🚀 Features

- Create and share text/code snippets easily
- RESTful API to store and retrieve pastes
- NGINX as a reverse proxy for routing
- Dockerized for easy deployment
- GitHub Actions for CI/CD pipeline automation
- Render hosting for live production
- Expiration feature for pastes

---

## 📦 Tech Stack

- Backend : Go (Golang)
- Frontend : HTML, CSS
- Web Server : NGINX
- CI/CD : GitHub Actions
- Containerization : Docker, Docker Compose
- Hosting : Render

---

## 📁 Project Structure

```
pastelite/
├── main.go                # Go backend server
├── public/
│   ├── index.html         # Home page (form input)
│   └── paste.html         # Page to view pastes
├── nginx/
│   └── default.conf       # NGINX reverse proxy config
├── Dockerfile             # Backend Docker image
├── docker-compose.yml     # Combined service stack
└── README.md
```

---

## 🐳 Getting Started (Local Development)

1. **Clone the repo**
```bash
git clone https://github.com/yourusername/pastelite.git
cd pastelite
```

2. **Run using Docker Compose**
```bash
docker-compose up --build
```

3. Open your browser and go to:  
[http://localhost:3000](http://localhost:3000)

---

## 🌐 Live Demo

[https://pastelite.onrender.com](https://pastelite.onrender.com)

---

## 📌 To-Do

- [ ] Syntax highlighting for code blocks
- [ ] Add dark mode theme
- [ ] Add support for private or unlisted pastes

---

## 👤 Author

**Tashu Choudhary**  
- [GitHub](https://github.com/TashuChoudhary)  
- [LinkedIn](https://linkedin.com/in/tashu-choudhary-30846227)

---