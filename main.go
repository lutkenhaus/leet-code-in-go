// Problem:
// - Given the heads of two sorted linked lists list1 and list2.
// Merge the lists into one sorted list.
// The list should be made by splicing together the nodes of the first two lists.
// - Return the head of the merged linked list.

// First thoughts:
// - At first I used a new mergedList and used append to populate this list.
// - Then I read the question again and saw that it asked to use splicing.
// - Splicing in this case basically means merging without creating new nodes.
// - Had to rewrite everything with this in mind.

// Approach:
// - Use splicing technique and an empty node as a head of the linked list.
// - Iterate through both lists until one is empty.
// - After the loop, point to the not empty list (if needed).

// Key:
// - The key to solving this problem is knowing what splicing is (I didn't at first).
// - Know how to navigate a linked list.
// - Use pointers efficiently.

// Complexity:
// - Time: O(n + m)
// - Space: O(1)

// Constraints:
// - The number of nodes in both lists is in the range [0, 50].
// - -100 <= Node.val <= 100
// - Both list1 and list2 are sorted in non-decreasing order.

// Optimizations:
// - The code is already at optimal performance O(n+m).
// - Removed some irrelevant functions.
// - Removed duplicated assignment inside the loop.

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

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
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

func (ll *LinkedList) AddIntegersToLinkedList(values []int) {
	if values == nil {
		return
	}
	if ll.Head == nil {
		ll.Head = NewNode(values[0])
	}

	if len(values) > 1 {
		current := ll.Head
		for i := 1; i < len(values); i++ {
			newListNode := NewNode(values[i])
			current.Next = newListNode
			current = current.Next
		}
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	aux := &ListNode{}
	current := aux

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return aux.Next
}

func PrintLinkedList(node *ListNode) {
	current := node
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {

	ll1 := NewEmptyList()
	ll2 := NewEmptyList()
	expectedValue1 := NewEmptyList()

	ll1.AddIntegersToLinkedList([]int{1, 2, 4})
	ll2.AddIntegersToLinkedList([]int{1, 3, 4})
	expectedValue1.AddIntegersToLinkedList([]int{1, 1, 2, 3, 4, 4})

	ll3 := NewEmptyList()
	ll4 := NewEmptyList()
	var expectedValue2 *ListNode = nil

	ll5 := NewEmptyList()
	ll6 := NewEmptyList()
	expectedValue3 := NewEmptyList()

	ll6.AddIntegersToLinkedList([]int{0})
	expectedValue3.AddIntegersToLinkedList([]int{0})

	testCases := []struct {
		value1        *ListNode
		value2        *ListNode
		expectedValue *ListNode
	}{
		{ll1.Head, ll2.Head, expectedValue1.Head},
		{ll3.Head, ll4.Head, expectedValue2},
		{ll5.Head, ll6.Head, expectedValue3.Head},
	}

	for i, tc := range testCases {
		mergedList := mergeTwoLists(tc.value1, tc.value2)
		expectedList := tc.expectedValue
		PrintLinkedList(mergedList)
		PrintLinkedList(expectedList)

		for expectedList != nil {
			if mergedList.Val != expectedList.Val {
				fmt.Printf("error: wanted (%v), got (%v)\n", mergedList.Val, expectedList.Val)
			}
			mergedList = mergedList.Next
			expectedList = expectedList.Next
		}
		fmt.Printf("\nTestcase: %d - Passed!\n\n", i+1)
	}
}
