// Problem: 14. Longest Common Prefix
// - Write a function to find the longest common prefix string amongst an array of strings.
// - If there is no common prefix, return an empty string "".

// First thoughts:
// - Does every word need to have the longest common prefix? I'll operate as if the answer is no.
// - Iterate through the words, saving common prefixes in a map.
// - My first solution was wrong, it did not solve every case.

// Approach:
// - I had to search for a solution.
// - The approach is vertical comparisons between letters.

// Key:
// - Know how to manipulate strings properly (which I didn't).
// - Vertical scanning algorithm.

// Complexity:
// - Time: O(n) -> n equals the sum of all characters (in all strings).
// - Space: O(1)

// Constraints:
// - (1 <= strs.length <= 200)
// - (0 <= strs[i].length <= 200)
// - strs[i] consists of only lowercase English letters if it is non-empty.

// Optimizations:
// - Vertical scanning is an optimal algorithm.

package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstWord := strs[0]
	for i := range firstWord {
		for j := 1; j < len(strs); j++ {
			currentWord := strs[j]
			if i >= len(currentWord) {
				return firstWord[:i]
			}
			if currentWord[i] != firstWord[i] {
				return firstWord[:i]
			}
		}
	}
	return firstWord
}

func main() {
	testCases := []struct {
		value         []string
		expectedValue string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dog"}, "dog"},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{}, ""},
		{[]string{"", "", "", ""}, ""},
		{[]string{"ghost", "gas", "guest", "gesture"}, "g"},
	}

	for i, tc := range testCases {
		response := longestCommonPrefix(tc.value)

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
