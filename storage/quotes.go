package storage

import (
	"math/rand"
)

var quotes = []string{
	"You have the right to work, but never to the fruit of work.",
	"A person can rise through the efforts of his own mind; or draw himself down, in the same manner.",
	"Be steadfast in yoga, O Arjuna. Perform your duty and abandon all attachment to success or failure.",
	"When meditation is mastered, the mind is unwavering like the flame of a lamp in a windless place.",
	"Set thy heart upon thy work, but never on its reward.",
	"For one who has conquered the mind, the mind is the best of friends.",
}

func GetRandomQuote() string {
	//rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}
