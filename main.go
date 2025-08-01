package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)
	zigzag := -1
	counter := 0

	for _, char := range s {
		rows[counter].WriteRune(char)
		if counter == 0 || counter == numRows-1 {
			zigzag *= -1
		}
		counter += zigzag
	}

	var result strings.Builder

	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
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
