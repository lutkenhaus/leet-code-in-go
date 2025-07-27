// Problem: 263. Ugly Number
// - An ugly number is a positive integer which does not have a prime factor other than
// 2, 3, and 5.
// - Given an integer n, return true if n is an ugly number.

// First thoughts:
// - Discard all negative numbers.
// - Divide by 2 until n can't be divided by 2 anymore.
// - Realized is not a good way to solve this.
// - Check if n is divisible by 5, 3 or 2.

// Approach:
// - Discard all negative numbers.
// - Test if n is divisible by either 5, 3 or 2, in that order.
// - If not, then return false.

// Key:
// - Make only necessary operations.

// Complexity:
// - Time: O(log n)
// - Space: O(1)

// Constraints:
// - (-2^31 <= n <= 2^31 - 1)

// Optimizations:
// - No optimizations so far.

package main

import (
	"fmt"
)

func uglyNumber(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%5 == 0 {
			n /= 5
		} else if n%3 == 0 {
			n /= 3
		} else if n%2 == 0 {
			n /= 2
		} else {
			return false
		}
	}
	return true
}

func main() {
	testCases := []struct {
		value         int
		expectedValue bool
	}{
		{6, true},
		{1, true},
		{14, false},
		{37, false},
		{49, false},
		{11, false},
		{83, false},
		{41, false},
		{-1, false},
		{-2, false},
		{-6, false},
	}

	for i, tc := range testCases {
		response := uglyNumber(tc.value)
		err := false

		if response != tc.expectedValue {
			fmt.Printf("Testcase %d error: wanted (%v), got (%v)\n", i+1, tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
