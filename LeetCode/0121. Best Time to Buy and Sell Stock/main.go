package main

import "fmt"

func maxProfit(prices []int) int {
	var profit = 0
	var min = prices[0]

	pricesLength := len(prices)

	for i := 0; i < pricesLength; i++ {
		if prices[i] < min {
			min = prices[i]
		} else if (prices[i] - min) > profit {
			profit = prices[i] - min
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	profit := maxProfit(prices)
	fmt.Println("Max profit: ", profit)
}
