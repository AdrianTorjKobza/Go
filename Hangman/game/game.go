package game

import (
	"fmt"
	"hangman/clear"
	"strings"
)

func Game(answer string, hangman []string, lives int, display []string) {
	// Start the game loop.
	for {
		// Print the current state of the hangman and the displayed word.
		fmt.Println(strings.Join(hangman, "\n"))
		fmt.Println(strings.Join(display, " "))

		// Prompt the player to guess a letter.
		fmt.Print("Guess a letter: ")
		var guess string
		fmt.Scanln(&guess)

		// Check if the guessed letter is in the answer.
		found := false

		for i, c := range answer {
			if string(c) == guess {
				display[i] = guess
				found = true
			}
		}

		// If the guessed letter is not in the answer, subtract a life.
		if !found {
			lives--
			if lives == 0 {
				// Game over.
				fmt.Println("You lose!")
				fmt.Println("The word was", answer)
				return
			}

			hangman[len(hangman)-lives-1] = hangman[len(hangman)-lives-1][:6] + "O" + hangman[len(hangman)-lives-1][7:]
		}

		// If the displayed word is equal to the answer, the player wins.
		if strings.Join(display, "") == answer {
			clear.ClearTerminal()
			fmt.Println(strings.Join(hangman, "\n"))
			fmt.Println(strings.Join(display, " "))
			fmt.Println("You win!")
			return
		}

		// Clear the terminal before the next guess.
		clear.ClearTerminal()
	}
}
