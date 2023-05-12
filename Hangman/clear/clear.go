package clear

import "fmt"

func ClearTerminal() {
	fmt.Print("\033[2J")   // Clear screen.
	fmt.Print("\033[H")    // Move cursor to top-left corner.
	fmt.Print("\033[?25l") // Hide cursor.
}
