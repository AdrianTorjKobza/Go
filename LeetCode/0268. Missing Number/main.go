package main

import "fmt"

func missingNumber(nums []int) int {
	var exist bool
	numsLength := len(nums)
	result := 0

	for i := 0; i <= numsLength; i++ {
		exist = false

		for _, value := range nums {
			if value == i {
				exist = true
			}
		}

		if !exist {
			result = i
		}
	}

	return result
}

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	missingNumber := missingNumber(nums)
	fmt.Println("Missing number: ", missingNumber) // Missing number: 8
}
