package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var sum int
	var output []int
	numsLength := len(nums)

	for i := 0; i < numsLength-1; i++ {
		for j := i + 1; j < numsLength; j++ {
			sum = nums[i] + nums[j]

			if sum == target {
				output = append(output, i, j)
			}
		}
	}

	return output
}

func main() {
	nums := []int{2, 7, 11, 15, 0}
	target := 11

	result := twoSum(nums, target)

	fmt.Println("Result:", result) //[2,4]
}
