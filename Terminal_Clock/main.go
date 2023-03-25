package main

import (
	"fmt"
	"time"
)

func main() {
	type clock_digits_type [5]string

	zero := clock_digits_type{
		"#####",
		"#   #",
		"#   #",
		"#   #",
		"#####",
	}

	one := clock_digits_type{
		" ### ",
		"   # ",
		"   # ",
		"   # ",
		"#####",
	}

	two := clock_digits_type{
		"#####",
		"    #",
		"#####",
		"#    ",
		"#####",
	}

	three := clock_digits_type{
		"#####",
		"    #",
		"#####",
		"    #",
		"#####",
	}

	four := clock_digits_type{
		"#   #",
		"#   #",
		"#####",
		"    #",
		"    #",
	}

	five := clock_digits_type{
		"#####",
		"#    ",
		"#####",
		"    #",
		"#####",
	}

	six := clock_digits_type{
		"#####",
		"#    ",
		"#####",
		"#   #",
		"#####",
	}

	seven := clock_digits_type{
		"#####",
		"    #",
		"    #",
		"    #",
		"    #",
	}

	eight := clock_digits_type{
		"#####",
		"#   #",
		"#####",
		"#   #",
		"#####",
	}

	nine := clock_digits_type{
		"#####",
		"#   #",
		"#####",
		"    #",
		"#####",
	}

	separator := clock_digits_type{
		"     ",
		"  *  ",
		"     ",
		"  *  ",
		"     ",
	}

	clock_digits := [...]clock_digits_type{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	rows := clock_digits[0] // Capture the number of string rows, used to create a number.

	// Infinite loop.
	for {
		fmt.Println("\x1b[2J") // Clear terminal screen macos.

		now := time.Now()                                        // Get the current time.
		hour, min, sec := now.Hour(), now.Minute(), now.Second() // Get the hour, min and sec.

		// Define the clock list of strings.
		clock := [...]clock_digits_type{
			clock_digits[hour/10], clock_digits[hour%10],
			separator,
			clock_digits[min/10], clock_digits[min%10],
			separator,
			clock_digits[sec/10], clock_digits[sec%10],
		}

		// Loop through the clock list and print the time in a single line.
		for r := range rows {
			for index, digit := range clock {
				next := clock[index][r]

				// Set emty string if the next digit is a separator and sec is divisible by 2.
				// The objective is to create the perception of a blinking clock.
				// This code will run every 2 sec.
				if digit == separator && sec%2 == 0 {
					next = "     "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second) // Add a delay of 1 sec, before displaying the clock.
	}
}
