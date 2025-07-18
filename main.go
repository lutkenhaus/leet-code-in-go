// Problem:
// - Given a sorted array of integers, in ascending order, remove the duplicates IN PLACE (no auxiliary array).
// - Order must remain.
// - Return the number of unique elements in the original array (nums).

// First thoughts:
// - Check constraints.
// - Since the solution needs to modify the original array, I could use 2 pointers.
// - Or use a HashMap to make the modifications.

// Approach:
// - My first approach was using a HashMap to store the unique numbers,
// while checking if the current number exists, and modifying the original nums array.
// - Works for tests cases 1 and 2.
// - After submitting the code passed for all 362 testcases.
// - I decided to also make a logic to use only 2 pointers.

// Key:
// - The key to solving this problem is to keep track of the position of the unique numbers,
// while jumping over the duplicates, making changes as needed.
// - I needed to search for a Two Pointer example, because I was overthinking the logic,
// it's simpler than I thought.

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

	uniqueNumberPointer := 1

	for i := 1; i < length; i++ {
		if nums[i] < minValue || nums[i] > maxValue {
			return -1
		}
		if nums[i] != nums[i-1] {
			nums[uniqueNumberPointer] = nums[i]
			uniqueNumberPointer++
		}
	}

	return uniqueNumberPointer
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
		{[]int{0, 0, 3, 4}, []int{0, 3, 4}},
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
