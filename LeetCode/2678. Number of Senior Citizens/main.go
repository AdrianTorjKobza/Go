package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	count := 0

	for _, value := range details {
		ageStr := value[11:13]
		ageInt, err := strconv.Atoi(ageStr)

		if err != nil {
			fmt.Println("Error converting age string to age int.", err)
		}

		if ageInt > 60 {
			count = count + 1
		}
	}

	return count
}

func main() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}

	countSeniors := countSeniors(details)
	fmt.Println("No. of seniors > 60 years old:", countSeniors) // No. of seniors > 60 years old: 2
}
