// Problem:
// - Given 2 strings A and B, return the index of the first occurrence of A in B.
// - Return -1 if not found.

// First thoughts:
// - Use the strings as an byte array. Then loop checking for matches.
// - Double loop, one for going through the haystack, when find possible match another loop.

// Approach:
// - Use byte arrays.
// - One loop to iterate the haystack, one loop to iterate the needle when possible match is found.

// Key:
// - Use an array of bytes/strings.
// - Search a substring inside a string.

// Constraints:
// - 1 <= a.length, b.length <= 104
// - a and b consist of only lowercase English characters.

// Optmizations:
// - The first solution worst case scenario was O(n x m). It could be improved.
// - New handlers for edge cases.
// - Found out that Knuth and Boyer-Moore algorithms are better to search for substrings.
// - Decided to improve my brute force solution anyway.
// - A new approach is to use string slicing. 3 lines of code, no nested loops.

package main

import (
	"fmt"
)

func strStr(needle string, haystack string) int {
	if needle == haystack || len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		needle        string
		haystack      string
		expectedValue int
	}{
		{"sad", "sadbutsad", 0},
		{"a", "abc", 0},
		{"sort", "somesortsome", 4},
		{"leetcode", "LeetCode", -1},
	}

	response := -1
	for _, tc := range testCases {
		response = strStr(tc.needle, tc.haystack)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		} else if response != -1 {
			fmt.Printf("%v in %v starts at position %d\n", tc.needle, tc.haystack, response)
		} else {
			fmt.Println("Needle is not in Haystack...")
		}
	}
}
