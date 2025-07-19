// Problem:
// - Given an integer array nums and an integer val, remove all ocurrences of val in nums IN-PLACE.
// The order of elements may be changed.
// Return the number of elements in nums which are not equal to val.
// The first k elements of nums must not be equal to val.

// First thoughts:
// - Check constraints.
// - Two pointer approach, similar to another leet-code problem I solved.
// - No auxiliary slice.
// - The input array is NOT sorted...

// Approach:
// - Two pointers.
// - I realized that the array needed to be sorted for the two pointers approach to work.
// - Used Go's sort.Int() built-in function.
// - Got stuck, had to search for a solution. It was in fact the two pointer solution all along.

// Key:
// - Use a two pointer approach, checking for every number that isn't the value to be removed.
// - No need to sort.

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length < 0 || length > 100 {
		return -1
	}
	if val < 0 || val > 100 {
		return -1
	}

	k := 0
	for i := 0; i < length; i++ {
		if nums[i] < 0 || nums[i] > 50 {
			return k
		}

		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
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
	fmt.Println("27. Remove Element")

	testCases := []struct {
		val          int
		nums         []int
		expectedNums []int
	}{
		{3, []int{3, 2, 2, 3}, []int{2, 2}},
		{2, []int{0, 1, 2, 2, 3, 0, 4, 2}, []int{0, 0, 1, 3, 4}},
	}

	for _, tc := range testCases {
		k := removeElement(tc.nums, tc.val)
		if k != len(tc.expectedNums) {
			return
		}
		sort.Ints(tc.nums[:k])
		for i, _ := range tc.expectedNums {
			if tc.nums[i] != tc.expectedNums[i] {
				return
			}
		}
		printSolution(tc.expectedNums)
		printSolution(tc.nums)
		fmt.Println()
	}
}
