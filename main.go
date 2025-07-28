package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	mapBrackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0, len(s)/2)

	for _, currentRune := range s {
		if value, exists := mapBrackets[currentRune]; exists {
			stack = append(stack, value)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != currentRune {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	testCases := []struct {
		value         string
		expectedValue bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([])", true},
		{"([)]", false},
	}

	for i, tc := range testCases {
		response := isValid(tc.value)

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
