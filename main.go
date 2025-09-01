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

func main() {
	testCases := []struct {
		value         *TreeNode
		expectedValue []int
	}{
		{&TreeNode{}, []int{1, 3, 4, 5, 7}},
	}

	for i, tc := range testCases {
		response := inorderTraversal(tc.value)

		err := false
		if !slices.Equal(response, tc.expectedValue) {
			fmt.Printf("Testcase %d error: wanted (%v), got (%v)\n", i+1, tc.expectedValue, response)
			err = true
		}
		if !err {
			fmt.Printf("Testcase %d OK!\n", i+1)
		}
	}
}
