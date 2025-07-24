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
	newListNode := &ListNode{Val: value}

	if ll.Head == nil {
		ll.Head = newListNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newListNode
}

func (ll *LinkedList) Preppend(value int) {
	// TODO: implement preppend logic
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	linkedList := LinkedList{}
	return linkedList.Head
}

func main() {
	testCases := []struct {
		value1        *ListNode
		value2        *ListNode
		expectedValue *ListNode
	}{
		{&ListNode{}, &ListNode{}, &ListNode{}},
		{&ListNode{}, &ListNode{}, &ListNode{}},
		{&ListNode{}, &ListNode{}, &ListNode{}},
	}

	response := &ListNode{}
	for i, tc := range testCases {
		response = mergeTwoLists(tc.value1, tc.value2)
		if response != tc.expectedValue {
			fmt.Printf("error: wanted (%v), got (%v)\n", tc.expectedValue, response)
		}
		fmt.Printf("Test passed for testcase: %d - a + b = %v\n", i, response)
	}
}
