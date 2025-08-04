package main

import (
	"fmt"
	"reflect"
)

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		nums[0] *= nums[0]
		return nums
	}

	result := make([]int, len(nums))
	resultP := len(result) - 1
	firstP := 0
	secondP := len(nums) - 1

	for firstP != secondP {
		if nums[firstP]*nums[firstP] >= nums[secondP]*nums[secondP] {
			result[resultP] = nums[firstP] * nums[firstP]
			firstP++
		} else {
			result[resultP] = nums[secondP] * nums[secondP]
			secondP--
		}
		resultP--
	}
	result[resultP] = nums[secondP] * nums[secondP]

	return result
}

func main() {
	testCases := []struct {
		value         []int
		expectedValue []int
	}{
		{[]int{0}, []int{0}},
		{[]int{0, 1}, []int{0, 1}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
		{[]int{-1, -2, 3, 4}, []int{1, 4, 9, 16}},
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for i, tc := range testCases {
		response := sortedSquares(tc.value)

		err := false
		if !reflect.DeepEqual(response, tc.expectedValue) {
			fmt.Printf("Testcase %d error: wanted (%v), got (%v)\n", i+1, tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
