package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}

	root.Left = sortedArrayToBST(nums[:middle])
	root.Right = sortedArrayToBST(nums[middle+1:])

	return root
}

func printLevelOrder(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}

	queue := []*TreeNode{root}
	result := []interface{}{}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				result = append(result, nil)
				continue
			} else {
				result = append(result, node.Val)
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	fmt.Println(result)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	testCases := []struct {
		value         []int
		expectedValue *TreeNode
	}{
		// Test case 1: [-10, -3, 0, 5, 9] -> [0, -3, 9, -10, null, 5]
		{
			value: []int{-10, -3, 0, 5, 9},
			expectedValue: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:  -3,
					Left: &TreeNode{Val: -10},
				},
				Right: &TreeNode{
					Val:  9,
					Left: &TreeNode{Val: 5},
				},
			},
		},
		// Test case 2: [1, 3] -> [3, 1] or [1, null, 3]
		{
			value: []int{1, 3},
			expectedValue: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
			},
		},
		// Test case 3: [] -> []
		{
			value:         []int{},
			expectedValue: nil,
		},
	}

	for i, tc := range testCases {
		response := sortedArrayToBST(tc.value)
		fmt.Printf("Expected response %d: ", i+1)
		printLevelOrder(response)
		if isSameTree(response, tc.expectedValue) {
			fmt.Printf("Test case %d OK!\n", i+1)
		} else {
			fmt.Printf("Test case %d error: wanted tree with root value %v, got tree with root value %v\n",
				i+1, tc.expectedValue, response)
		}
	}
}
