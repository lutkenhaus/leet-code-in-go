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
// -

package main

import (
	"fmt"
)

func strStr(haystack, needle string) int {
	h := []byte(haystack)
	n := []byte(needle)
	counter := 0

	for i, haystackValue := range h {
		counter = 0
		if haystackValue == n[0] {
			for z, needleValue := range n {
				if i+z >= len(h) {
					break
				}
				if needleValue != h[i+z] {
					break
				} else {
					counter++
					if counter == len(n) {
						return i
					}
				}
			}
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
		response = strStr(tc.haystack, tc.needle)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		if response != -1 {
			fmt.Printf("%v in %v starts at position %d\n", tc.needle, tc.haystack, response)
		} else {
			fmt.Println("Needle is not in Haystack...")
		}
	}
}
