package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		sum            map[string]int
		events         []string // Slice to store the unique events name.
		total_duration int      // Store the total duration of all events.
		line           int      // Store the line where the error occured.
	)

	sum = make(map[string]int) // Initialize the map, to store all events, from the log file.

	log := bufio.NewScanner(os.Stdin)

	// Scan the log file, line by line.
	for log.Scan() {
		line++

		text_columns := strings.Fields(log.Text())

		// Check if the format of the log file is correct.
		if len(text_columns) != 2 {
			fmt.Printf("Wrong input: %v column(s). [line #%d] \nWe are expecting a log file with two columns.\n", len(text_columns), line)
			return
		}

		event := text_columns[0]                         // Capture the Event column, from log file.
		duration, error := strconv.Atoi(text_columns[1]) // Capture the Duration column, from the log file.

		// Check if the input is correct.
		if duration < 0 || error != nil {
			fmt.Printf("Wrong input: %q [line #%d] \nWe are expecting a positive number for Duration.\n", text_columns[1], line)
			return
		}

		// Check if the event is in the map.
		if _, ok := sum[event]; !ok {
			events = append(events, event) // If event doesn't exist in the map, add it to the slice.
		}

		total_duration = total_duration + duration
		sum[event] = sum[event] + duration

	}

	// Print the two columns content, from log file.
	fmt.Printf("%-30s %10s\n", "EVENT", "DURATION (ms)")
	fmt.Println(strings.Repeat("-", 45))

	for _, event := range events {
		duration := sum[event]
		fmt.Printf("%-30s %10d\n", event, duration)
	}

	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-30s %10d\n", "TOTAL", total_duration)
}
