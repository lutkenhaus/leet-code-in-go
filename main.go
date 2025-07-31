package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}

	if numRows == 1 || numRows == len(s) {
		return s
	}

	mapMatrix := make(map[int]string, numRows)
	zigzag := false
	counter := 0

	for _, v := range s {
		switch {
		case zigzag:
			mapMatrix[counter] += string(v)
			if counter <= 0 {
				zigzag = false
				counter++
			} else {
				counter--
			}
		case !zigzag:
			mapMatrix[counter] += string(v)
			if counter >= numRows-1 {
				zigzag = true
				counter--
			} else {
				counter++
			}
		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		result += mapMatrix[i]
	}

	return result
}

func main() {
	testCases := []struct {
		value         string
		numRows       int
		expectedValue string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"", 1, ""},
		{"string", 1, "string"},
		{"PAYPAL", 6, "PAYPAL"},
		{"PAYPAL", 2, "PYAAPL"},
	}

	for i, tc := range testCases {
		response := convert(tc.value, tc.numRows)

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
