// Problem:
// - Given an integer x, return true if x is a palindrome, and false otherwise.

// First thoughts:
// - Create a reversed copy of x.
// - Compare the copy of x to x.

// Approach:
// - Transform the interger into an array on integers.
// - Reverse the integer using a inverted loop.
// - Changed approach to use math in order to invert the value.
// - Kept two slices, one in original order, other one inverted.
// - Compare two slices to check if x is a palindrome.

// Key:
// - Use math. Remainder of x divided by 10 is equal to last digit.
// - remainder = x % 10
// - Use two slices for comparison.

// Constraints:
// - ( -231 <= x <= 231 - 1 )

// Optmizations:
// - No need to use slices for this problem, only realized this after looking up other solutions.
// - Math is powerful.

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}

	if x < 0 {
		return false
	}

	originalX := x
	reversedX := 0

	for x > 0 {
		reversedX = reversedX*10 + x%10
		x /= 10
	}

	return originalX == reversedX
}

func main() {
	fmt.Println("9. Palindrome Number")
	testCases := []struct {
		value         int
		expectedValue bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{542, false},
		{777, true},
	}

	response := false
	for i, tc := range testCases {
		response = isPalindrome(tc.value)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		fmt.Printf("PASSED for test case (%d): %v - %v\n", i, tc.value, response)
	}
}
