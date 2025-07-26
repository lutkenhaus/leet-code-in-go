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

		err := false
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
