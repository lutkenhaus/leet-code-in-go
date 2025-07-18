// Problem:
// - Given a sorted array of integers, in ascending order, remove the duplicates IN PLACE (no auxiliary array).
// - Order must remain.
// - Return the number of unique elements in the original array (nums).

// First thoughts:
// - Check constraints.
// - Since the solution needs to modify the original array, I could use 2 pointers.
// - Or use a HashMap to make the modifications.

// Approach:
// -

// Key:

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxLength = 3e4
	minValue  = -100
	maxValue  = 100
)

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 1 || length > int(maxLength) {
		return -1
	}

	counter := 0
	uniqueNumbers := make(map[int]int)
	for index, value := range nums {
		if nums[index] < minValue || nums[index] > maxValue {
			return -1
		}

		if _, exists := uniqueNumbers[value]; !exists {
			uniqueNumbers[value] = index
			nums[counter] = value
			counter++
		}
	}

	return counter
}

func printSolution(nums []int) {
	if len(nums) == 0 {
		fmt.Println(nums)
	}
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	result := "[" + strings.Join(strNums, ", ") + "]"
	fmt.Println(result)
}

func main() {
	fmt.Println("26. Remove Duplicates from Sorted Array")

	testCases := []struct {
		nums         []int
		expectedNums []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{3, 3}, []int{3}},
		{[]int{0, 3, 4, 0}, []int{0, 3, 4}},
		{[]int{-3, 3, 4, 90}, []int{-3, 3, 4, 90}},
	}

	for _, tc := range testCases {
		k := removeDuplicates(tc.nums)
		if k != len(tc.expectedNums) {
			fmt.Println("error here")
			return
		}
		for i, _ := range tc.expectedNums {
			if tc.nums[i] != tc.expectedNums[i] {
				return
			}
		}
		printSolution(tc.expectedNums)
	}
}
