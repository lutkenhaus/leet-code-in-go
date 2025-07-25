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
// - The number of nodes in both lists is in the range [0, 50].
// - -100 <= Node.val <= 100
// - Both list1 and list2 are sorted in non-decreasing order.

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
	Val  *int
	Next *ListNode
}

func NewEmptyNode() *ListNode {
	return &ListNode{
		Val:  nil,
		Next: nil,
	}
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  &val,
		Next: nil,
	}
}

type LinkedList struct {
	Head *ListNode
}

func NewEmptyList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func NewLinkedList(node ListNode) *LinkedList {
	return &LinkedList{
		Head: &node,
	}
}

func (ll *LinkedList) Append(value *int) {
	if value == nil {
		fmt.Printf("possible nil pointer dereference error")
		return
	}

	newListNode := NewNode(*value)

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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	mergedLinkedList := NewEmptyList()
	var currentListNode1, currentListNode2 *ListNode
	if list1 != nil {
		if list1.Val != nil && list1.Next != nil {
			currentListNode1 = &ListNode{Val: list1.Val, Next: list1.Next}
		} else if list1.Val != nil {
			currentListNode1 = &ListNode{Val: list1.Val}
		}
	}
	if list2 != nil {
		if list2.Val != nil && list2.Next != nil {
			currentListNode2 = &ListNode{Val: list2.Val, Next: list2.Next}
		} else if list2.Val != nil {
			currentListNode2 = &ListNode{Val: list2.Val}
		}
	}

	for currentListNode1 != nil || currentListNode2 != nil {
		switch {
		case currentListNode1 == nil:
			mergedLinkedList.Append(currentListNode2.Val)
			currentListNode2 = currentListNode2.Next
		case currentListNode2 == nil:
			mergedLinkedList.Append(currentListNode1.Val)
			currentListNode1 = currentListNode1.Next
		case *currentListNode1.Val > *currentListNode2.Val:
			mergedLinkedList.Append(currentListNode2.Val)
			currentListNode2 = currentListNode2.Next
		case *currentListNode1.Val < *currentListNode2.Val:
			mergedLinkedList.Append(currentListNode1.Val)
			currentListNode1 = currentListNode1.Next
		default:
			mergedLinkedList.Append(currentListNode1.Val)
			mergedLinkedList.Append(currentListNode2.Val)
			currentListNode1 = currentListNode1.Next
			currentListNode2 = currentListNode2.Next
		}
	}

	fmt.Println("")
	mergedLinkedList.PrintLinkedList()

	return mergedLinkedList.Head
}

func (ll *LinkedList) PrintLinkedList() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", *current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {

	zero := 0
	one := 1
	two := 2
	three := 3
	four := 4

	ll1 := NewEmptyList()
	ll2 := NewEmptyList()
	expectedValue1 := NewEmptyList()

	ll1.Append(&one)
	ll1.Append(&two)
	ll1.Append(&four)

	ll2.Append(&one)
	ll2.Append(&three)
	ll2.Append(&four)

	expectedValue1.Append(&one)
	expectedValue1.Append(&one)
	expectedValue1.Append(&two)
	expectedValue1.Append(&three)
	expectedValue1.Append(&four)
	expectedValue1.Append(&four)

	ll3 := NewEmptyList()
	ll4 := NewEmptyList()
	expectedValue2 := NewEmptyNode()

	ll5 := NewEmptyList()
	ll6 := NewEmptyList()
	expectedValue3 := NewEmptyList()

	ll6.Append(&zero)
	expectedValue3.Append(&zero)

	testCases := []struct {
		value1        *ListNode
		value2        *ListNode
		expectedValue *ListNode
	}{
		{ll1.Head, ll2.Head, expectedValue1.Head},
		{ll3.Head, ll4.Head, expectedValue2},
		{ll5.Head, ll6.Head, expectedValue3.Head},
	}

	response := &ListNode{}
	for i, tc := range testCases {
		currentMergedList := &ListNode{}
		currentExpectedList := &ListNode{}
		response = mergeTwoLists(tc.value1, tc.value2)

		// TODO: implement better test cases check...
		switch {
		case tc.value1 == nil || tc.value2 == nil:
			fmt.Printf("possible empty lists, use debugger to check the return value")
		default:
			currentMergedList = response
			currentExpectedList = tc.expectedValue
			for currentExpectedList != nil {
				if *currentMergedList.Val != *currentExpectedList.Val {
					fmt.Printf("error: wanted (%v), got (%v)\n", *currentMergedList.Val, *currentExpectedList.Val)
				}
				currentMergedList = currentMergedList.Next
				currentExpectedList = currentExpectedList.Next
			}
			fmt.Printf("Testcase: %d - Passed!\n", i)
		}
	}
}
