package main

import (
	"fmt"
	"strings"
)

func someFunc(s string) int {
	trimmed := strings.TrimSpace(s)

	counter := 0
	for i := len(trimmed) - 1; i >= 0; i-- {
		if trimmed[i] == ' ' {
			return counter
		}
		counter++
	}
	return counter
}

func main() {
	testCases := []struct {
		value         string
		expectedValue int
	}{
		{"", 0},
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"test", 4},
	}

	for i, tc := range testCases {
		response := someFunc(tc.value)

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
