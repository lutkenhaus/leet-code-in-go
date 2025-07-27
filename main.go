// Problem: 14. Longest Common Prefix
// - Write a function to find the longest common prefix string amongst an array of strings.
// - If there is no common prefix, return an empty string "".

// First thoughts:
// - Does every word need to have the longest common prefix? I'll operate as if the answer is no.
// - Iterate through the words, saving common prefixes in a map.

// Approach:
// -

// Key:
// -

// Complexity:
// -

// Constraints:
// - (1 <= strs.length <= 200)
// - (0 <= strs[i].length <= 200)
// - strs[i] consists of only lowercase English letters if it is non-empty.

// Optimizations:
// -

package main

import (
	"fmt"
)

func someFunc(strs []string) string {
	strMap := make(map[string]int)
	for _, word := range strs {
		prefix := string(word[0])
		for z := 0; z < len(word)-1; z++ {
			strMap[prefix]++
			if z < len(word)-1 {
				prefix = prefix + string(word[z+1])
			}
		}
	}

	longestCommonPrefix := ""
	counter := 0
	for i, value := range strMap {
		if value >= counter {
			longestCommonPrefix = i
			counter = value
		}
	}

	if counter > 1 {
		return longestCommonPrefix
	}
	return ""
}

func main() {
	testCases := []struct {
		value         []string
		expectedValue string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ghost", "gas", "guest", "gesture"}, "g"},
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
