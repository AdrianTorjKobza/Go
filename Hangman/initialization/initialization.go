package initialization

import (
	"math/rand"
)

// Define the list of words to choose from.
var words = []string{"apple", "banana", "cherry", "orange", "pineapple"}

// Choose a random word from the list
func RandomWord() string {
	answer := words[rand.Intn(len(words))]
	return answer
}

// Create a slice of underscores with the same length as the answer.
func Underscores(answer string) []string {
	display := make([]string, len(answer))

	for i := range display {
		display[i] = "_"
	}

	return display
}

// Initialize the hangman and the number of lives.
func InitHangman() ([]string, int) {
	hangman := []string{
		"  +---+",
		"  |   |",
		"      |",
		"      |",
		"      |",
		"      |",
		"=========",
	}

	return hangman, len(hangman)
}
