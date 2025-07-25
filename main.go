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
	Val  int
	Next *ListNode
}

func NewEmptyNode() *ListNode {
	return &ListNode{}
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (ln *ListNode) IsEmpty() bool {
	if ln == nil {
		return true
	}

	if ln.Val == 0 && ln.Next == nil {
		return true
	}

	return false
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
		return &ListNode{}
	}
	aux := &ListNode{}
	current := aux

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			current = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			current = list2
			list2 = list2.Next
		}
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
	expectedValue2 := NewEmptyNode()

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

	response := &ListNode{}
	for i, tc := range testCases {
		currentMergedList := &ListNode{}
		currentExpectedList := &ListNode{}
		response = mergeTwoLists(tc.value1, tc.value2)

		if tc.expectedValue == nil && response == nil {
			fmt.Printf("empty lists")
			fmt.Printf("Testcase: %d - Passed!\n", i)
		}

		currentMergedList = response
		currentExpectedList = tc.expectedValue
		fmt.Printf("\nmerged list: ")
		PrintLinkedList(currentMergedList)
		fmt.Printf("\nexpected list: ")
		PrintLinkedList(currentExpectedList)
		for currentExpectedList != nil {
			if currentMergedList.Val != currentExpectedList.Val {
				fmt.Printf("error: wanted (%v), got (%v)\n", currentMergedList.Val, currentExpectedList.Val)
			}
			currentMergedList = currentMergedList.Next
			currentExpectedList = currentExpectedList.Next
		}
		fmt.Printf("Testcase: %d - Passed!\n", i)

	}
}
