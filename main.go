package main

import (
	"fmt"
)

// Problem:
// - return 2 indices of the 2 numbers that add upt to the target.
// - there is only one solution.
// - cannot use the same element twice.

// First thoughts:
// - create logic for constraints.
// - use an auxiliar slice with all the numbers where their value is lesser than the target.
// - test all the possibilities using the smaller slice.

// Approach:
// - the filter approach was no good, since I didn't account for the fact that
// the numbers can be negative or zero.
// - brute force is an option but is not optimal.

// Key:
// - the key to solving this problem in an efficient way is to use a HashMap.
// - I had to search for tips and solutions in order to figure this out.
// - storing the [value] as the index of the HashMap for every visited number in the array
// allows us to check if the complement (target - number) of the current number has already been seen.

const (
	maxLength = 1e4
	minValue  = -1e9
	maxValue  = 1e9
)

func twoSum(nums []int, target int) []int {
	lenArr := len(nums)

	if lenArr < 2 || lenArr > int(maxLength) {
		return []int{}
	}
	if target < int(minValue) || target > int(maxValue) {
		return []int{}
	}

	seenNumbers := make(map[int]int)

	for i, n := range nums {
		if n < int(minValue) || n > int(maxValue) {
			return []int{}
		}
		complement := target - n
		if j, exists := seenNumbers[complement]; exists {
			return []int{j, i}
		}
		seenNumbers[n] = i
	}
	return []int{}
}

func printSolution(res []int) {
	if len(res) != 2 {
		fmt.Println("Invalid result length")
		return
	}
	fmt.Printf("[%d, %d]\n", res[0], res[1])
}

func main() {
	fmt.Println("1. Two Sum")

	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
		{[]int{0, 4, 3, 0}, 0},
		{[]int{-3, 4, 3, 90}, 0},
	}

	for _, tc := range testCases {
		res := twoSum(tc.nums, tc.target)
		printSolution(res)
	}
}
