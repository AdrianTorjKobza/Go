package main

import (
	"hangman/game"
	"hangman/initialization"
)

func main() {
	answer := initialization.RandomWord()
	hangman, lives := initialization.InitHangman()
	display := initialization.Underscores(answer)
	game.Game(answer, hangman, lives, display)
}
