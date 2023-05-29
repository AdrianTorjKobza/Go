package main

import "fmt"

func lengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				return count
			}
		} else {
			count++
		}
	}

	return count
}

func main() {
	stringInput := "   fly me   to   the moon  "

	wordLength := lengthOfLastWord(stringInput)
	fmt.Println("Length of the last word: ", wordLength) // Length of the last word:  4
}
