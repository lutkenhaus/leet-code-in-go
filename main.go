// Problem:
// - Given the heads of two sorted linked lists list1 and list2.
// Merge the lists into one sorted list.
// The list should be made by splicing together the nodes of the first two lists.
// - Return the head of the merged linked list.

// First thoughts:
// -

// Approach:
// -

// Key:
// -

// Complexity:
// -

// Constraints:
// -

// Optimizations:
// -

package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

func (ll *LinkedList) Append(value int) {
	// TODO: create a logic to append a value to the list.
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	linkedList := LinkedList{}
	return linkedList.Head
}

func main() {
	testCases := []struct {
		value         string
		expectedValue string
	}{
		{"LeetCode", "LeetCode!"},
		{"LeetCode!", "LeetCode!!"},
		{"LeetCode!!", "LeetCode!!!"},
	}

	response := ""
	for i, tc := range testCases {
		response = someFunc(tc.value)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		fmt.Printf("Test passed for testcase: %d - a + b = %v\n", i, response)
	}
}
