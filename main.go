// Problem:
// - Write an algorithm to determine if a number n is happy.
// - A happy number is a number defined by the following process:
// - - Starting with any positive integer, replace the number by the sum of the squares of its digits.
// - - Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// - - Those numbers for which this process ends in 1 are happy.
// - Return true if n is a happy number, and false if not.

// First thoughts:
// - Is there a way to know before hand if a number won't be happy?
// - When to stop the loop?

// Approach:
// - Try 30 times, if the number is not happy then it is depressed (stop infinite loop).
// - Use math to extract the digits from n, and an auxiliary variable to store the sum.
// - Stops if either n is 1 or counter is 0.

// Key:
// - Use the modulus to extract the digits from n.
// - Loop until n is happy (or depressed).

// Complexity:
// -

// Constraints:
// - 1 <= n <= 2^31 - 1

// Optimizations:
// -

package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
	counter := 30
	sum := 0
	if n == 1 {
		return true
	}

	for {
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		if sum == 1 {
			return true
		}
		n = sum
		sum = 0

		if counter == 0 {
			return false
		}
		counter--
	}
}

func main() {
	testCases := []struct {
		value         int
		expectedValue bool
	}{
		{19, true},
		{2, false},
		{1, true},
		{int(math.Pow(2, 30)), false},
		{int(math.Pow(2, 31)), false},
		{int(math.Pow(2, 99)), false},
		{1002, false},
		{989, true},
	}

	for i, tc := range testCases {
		response := isHappy(tc.value)

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
