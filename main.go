package main

import (
	"fmt"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveTraversal(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	recursiveTraversal(node.Left, result)
	*result = append(*result, node.Val)
	recursiveTraversal(node.Right, result)
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	recursiveTraversal(root, &result)

	return result
}

// toInt creates a pointer to an int for use in the input slice.
func toInt(val int) *int {
	return &val
}

// createTreeFromSlice constructs a binary tree from a level-order slice of nullable integers.
// The input slice represents a breadth-first traversal, where nil indicates an absent node.
// Returns the root of the tree or nil if the input is empty or invalid.
func createTreeFromSlice(arr []*int) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	// Create root node
	root := &TreeNode{Val: *arr[0]}
	nodes := make([]*TreeNode, 0, len(arr)/2+1) // Pre-allocate queue with estimated capacity
	nodes = append(nodes, root)
	front := 0 // Queue front index
	index := 1 // Input array index

	for front < len(nodes) && index < len(arr) {
		current := nodes[front]
		front++

		// Process left child
		if index < len(arr) && arr[index] != nil {
			current.Left = &TreeNode{Val: *arr[index]}
			nodes = append(nodes, current.Left)
		}
		index++

		// Process right child
		if index < len(arr) && arr[index] != nil {
			current.Right = &TreeNode{Val: *arr[index]}
			nodes = append(nodes, current.Right)
		}
		index++
	}

	return root
}

func inOrderTraversal(root *TreeNode) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Val)
		inOrderTraversal(root.Right)
	}
}

func main() {
	testCases := []struct {
		name          string
		value         []*int
		expectedValue []int
	}{
		{
			name:          "Empty tree",
			value:         []*int{},
			expectedValue: []int{},
		},
		{
			name:          "Single node",
			value:         []*int{toInt(1)},
			expectedValue: []int{1},
		},
		{
			name:          "Left subtree only",
			value:         []*int{toInt(5), toInt(3), nil, toInt(1)},
			expectedValue: []int{1, 3, 5},
		},
		{
			name:          "Complete binary tree",
			value:         []*int{toInt(4), toInt(2), toInt(6), toInt(1), toInt(3), toInt(5), toInt(7)},
			expectedValue: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:          "Unbalanced tree",
			value:         []*int{toInt(5), toInt(3), toInt(7), toInt(1), nil, toInt(6)},
			expectedValue: []int{1, 3, 5, 6, 7},
		},
		{
			name:          "Tree with negative values",
			value:         []*int{toInt(0), toInt(-5), toInt(5)},
			expectedValue: []int{-5, 0, 5},
		},
	}

	for i, tc := range testCases {
		root := createTreeFromSlice(tc.value)
		response := inorderTraversal(root)

		err := false
		if !slices.Equal(response, tc.expectedValue) {
			fmt.Printf("Testcase %d error: wanted (%v), got (%v)\n", i+1, tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Println("In-order traversal of the created tree:")
			inOrderTraversal(root)
			fmt.Println()
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
