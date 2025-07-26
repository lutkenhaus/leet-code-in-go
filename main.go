// Problem:
// - Given a large integer represented as an integer array (digits := [1, 2, 3, ...]), where
// each digits[i] is the ith digit of the integer.
// - The digits are ordered from most significant to least significant in left-to-right order.
// - The large integer does not contain any leading 0's.
// - Increment the large integer by one and return the resulting array of digits.

// First thoughts:
// - This seems easy, way to easy it got me suspicious.
// - Add 1 to digits[len(digits) - 1].
// - Solve edge cases.

// Approach:
// - Check if last digit is not 9, then add 1 to the last position of the array.
// - If it is 9 then change the last digit to 1, and append 0.

// Key:
// - Understand that there is no need to use type conversions ([]int to int).
// - Remember edge cases.

// Complexity:
// -

// Constraints:
// - 1 <= digits.length <= 100
// - 0 <= digits[i] <= 9
// - digits does not contain any leading 0's.

// Optimizations:
// -

package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	arrLen := len(digits)
	if digits[arrLen-1] != 9 {
		digits[arrLen-1] += 1
	} else {
		digits[arrLen-1] = 1
		digits = append(digits, 0)
	}

	return digits
}

func main() {
	testCases := []struct {
		value         []int
		expectedValue []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for i, tc := range testCases {
		response := plusOne(tc.value)
		for z, v := range tc.expectedValue {
			if response[z] != v {
				fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
			}
		}
		fmt.Printf("\n%+v\n%+v", response, tc.expectedValue)
		fmt.Printf("\nTestcase %d passed!\n", i+1)
	}
}
