package utils

import (
	"math/rand"
	"time"
)


// Function returns a random quote from an array with quotes of type string
// It generates a random index using randomInteger function and
// returns a quote at the random index
func RandomQuote(quotes []string) string {
	randomIndex := randomInteger(0, len(quotes) - 1)

	quote := quotes[randomIndex]

	return quote
}

// Helper function to generate a random integer in a given range
// Using Unix time to generate a new seed
// which guarantees that the number returned will be truly random
func randomInteger(min, max int) int {

	seed := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(seed)

	randomNumber := generator.Intn(max - min) + min

	return randomNumber
}