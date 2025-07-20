// Problem:
// - Given a roman numeral (string), convert it to an integer.

// First thoughts:
// - Create a map to represent each symbol and it's correspondent value.
// - Convert the input string to a string slice.
// - Look at leetcode hints.
// - Only one hint, work at the string backwards.
// - The rules for roman numbers are:
// - - If a smaller character appears before a greater character, it's a subtraction.
// - - Equal or smaller letters to the right are summed.
// - - Same character can be summed only 3 times.

// Approach:
// - First I made 3 variables, current, next and previous. At every iteration I was trying to thinking of
// a logic to do comparisons, checking for relevant letters around the current character.
// - Got nowhere.
// - Then I stopped trying to loop backwards, and searched for a solution online.
// - After reading a solution it became obvious that only 2 operations matter, addition and subtraction.

// Key:
// - Use the string as a slice of strings (or bytes).
// - Check if the current character is smaller (subtract), equal or greater (add) than the next character.
// - Use a variable to store the total sum.

package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	mapSymbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	strSlice := strings.Split(s, "")
	length := len(strSlice)

	totalSum := 0
	for i := 0; i < length; i++ {
		if i < length-1 && mapSymbols[strSlice[i]] < mapSymbols[strSlice[i+1]] {
			totalSum -= mapSymbols[strSlice[i]]
		} else {
			totalSum += mapSymbols[strSlice[i]]
		}
	}

	return totalSum
}

func main() {
	fmt.Println("13. Roman to Integer")

	testCases := []struct {
		characters    string
		expectedValue int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MCMXCIX", 1999},
		{"MMXXV", 2025},
		{"MM", 2000},
	}

	for _, tc := range testCases {
		response := romanToInt(tc.characters)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted %d, got (%v)\n", tc.expectedValue, response)
			return
		}
		fmt.Printf("- %v = %d\n", tc.characters, response)
	}
}
