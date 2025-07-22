// Problem:
// -

// First thoughts:
// -

// Approach:
// -

// Key:
// -

// Complexity:
// -

// Constraints:
// -

// Optimizations:
// -

package main

import (
	"fmt"
)

func someFunc(s string) string {
	return s + "!"
}

func main() {
	testCases := []struct {
		value         string
		expectedValue string
	}{
		{"LeetCode", "LeetCode!"},
		{"LeetCode!", "LeetCode!!"},
		{"LeetCode!!", "LeetCode!!!"},
	}

	response := ""
	for i, tc := range testCases {
		response = someFunc(tc.value)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		fmt.Printf("Test passed for testcase: %d - a + b = %v\n", i, response)
	}
}
