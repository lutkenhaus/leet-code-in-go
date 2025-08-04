package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if n == 0 {
		return nums1
	}

	firstP := m - 1
	secondP := n - 1
	thirdP := m + n - 1

	for firstP >= 0 && secondP >= 0 {
		if nums1[firstP] >= nums2[secondP] {
			nums1[thirdP] = nums1[firstP]
			firstP--
		} else {
			nums1[thirdP] = nums2[secondP]
			secondP--
		}
		thirdP--
	}

	for secondP >= 0 {
		nums1[thirdP] = nums2[secondP]
		secondP--
		thirdP--
	}

	return nums1
}

func main() {
	testCases := []struct {
		nums1         []int
		m             int
		nums2         []int
		n             int
		expectedValue []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for i, tc := range testCases {
		response := merge(tc.nums1, tc.m, tc.nums2, tc.n)

		err := false
		if !cmp.Equal(response, tc.expectedValue) {
			fmt.Printf("Testcase %d error: wanted (%v), got (%v)\n", i+1, tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
