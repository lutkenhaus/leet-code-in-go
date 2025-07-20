// Problem:
// -

// First thoughts:
// -

// Approach:
// -

// Key:
// -

// Constraints:
// -

// Optmizations:
// -

package main

import (
	"fmt"
)

func someFunc(string) string {
	return "LeetCode!"
}

func main() {
	testCases := []struct {
		value         string
		expectedValue string
	}{
		{"LeetCode!", "LeetCode!"},
		{"LeetCode", "LeetCode!"},
		{"leetcode", "LeetCode!"},
	}

	for _, tc := range testCases {
		response := someFunc(tc.value)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted %v, got (%v)\n", tc.expectedValue, response)
			return
		}
		fmt.Printf("%v - %v\n", tc.value, response)
	}
}
