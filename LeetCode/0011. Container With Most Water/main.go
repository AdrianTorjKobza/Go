package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxArea := 0
	area := 0
	minHeight := 0
	left := 0
	right := len(height) - 1

	for left < right {
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		width := (right - left)
		area = minHeight * width

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	area := maxArea(height)
	fmt.Println("Max amount of water:", area) // Max amount of water: 49
}
