// Problem:
// - Fiven two binary strings a and b, return their sum as a binary string.

// First thoughts:
// - Write an algorithm that mimics step by step the math operation.
// - Use string indexation supported by Go.
// - Is it possible to solve using two pointers?
// - Think about rules of binary sum.

// Approach:
// - Created an algorithm that copies the same steps as the mathematics approach to binary sum.
// - Used the strings as byte slices unneficientlyn (and regretted it deeply after many unnecessary shenanigans).
// - Kept in mind all of the binary sum rules.

// Key:
// - Understand binary sum (every step).
// - Know how to manipulate strings efficiently.

// Constraints:
// - 1 <= a.length, b.length <= 104
// - a and b consist only of '0' or '1' characters.
// - Each string does not contain leading zeros except for the zero itself.

// Optimizations:
// - Space complexity can be improved.
// - The only auxiliary variable needed is the result (or response). No need to use any other to store values.
// - It's possible to process both strings from right to left with one loop.
// - Byte handling can be more efficient, especially if the % and the / operands are used to store the result and carry

package main

import (
	"fmt"
)

func addBinary(a, b string) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	if a == "0" && b == "0" {
		return "0"
	}
	if a == "0" && b == "1" || a == "1" && b == "0" {
		return "1"
	}

	result := ""
	carry := 0
	lenA := len(a)
	lenB := len(b)
	commonLength := lenA

	if lenA != lenB {
		if lenA > lenB {
			j := lenB - 1
			bytesB := make([]byte, lenA)
			for i := lenA - 1; i >= lenA-lenB; i-- {
				bytesB[i] = b[j] - '0'
				j--
			}
			strB := ""
			for _, v := range bytesB {
				strB = fmt.Sprintf("%s%v", strB, v)
			}
			b = strB
		} else {
			j := lenA - 1
			bytesA := make([]byte, lenB)
			for i := lenB - 1; i >= lenB-lenA; i-- {
				bytesA[i] = a[j] - '0'
				j--
			}
			strA := ""
			for _, v := range bytesA {
				strA = fmt.Sprintf("%s%v", strA, v)
			}
			a = strA
			commonLength = lenB
		}
	}

	for i := commonLength - 1; i >= 0; i-- {
		aBit := int(a[i] - '0')
		bBit := int(b[i] - '0')

		switch {
		case carry == 0:
			sum := aBit + bBit
			if sum == 0 {
				result = "0" + result
			}
			if sum == 1 {
				result = "1" + result
			} else if sum == 2 {
				result = "0" + result
				carry = 1
			}
		case carry == 1:
			sum := aBit + bBit + carry
			if sum == 1 {
				result = "1" + result
				carry = 0
			} else if sum == 2 {
				result = "0" + result
				carry = 1
			} else if sum == 3 {
				result = "1" + result
				carry = 1
			}
		default:
			continue
		}
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}

func main() {
	testCases := []struct {
		a             string
		b             string
		expectedValue string
	}{
		{"10111", "10", "11001"},
		{"10", "1011100", "1011110"},
		{"101010", "10", "101100"},
		{"1010", "1011", "10101"},
		{"11", "1", "100"},
		{"1", "1", "10"},
		{"0", "0", "0"},
	}

	response := ""
	for i, tc := range testCases {
		response = addBinary(tc.a, tc.b)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		fmt.Printf("Test passed for testcase: %d - a + b = %v\n", i, response)
	}
}
