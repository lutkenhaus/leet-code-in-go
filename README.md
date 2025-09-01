# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 94. Binary Tree Inorder Traversal
- Given the root of a binary tree, return the inorder traversal of its nodes' values.

- Example 1:

Input: root = [1,null,2,3]
Output: [1,3,2]

- Example 2:
Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
Output: [4,2,6,5,7,1,3,9,8]

- Example 3:
Input: root = []
Output: []

- Example 4:
Input: root = [1]
Output: [1]

Follow up: Recursive solution is trivial, could you do it iteratively?

# First thoughts:
- Figure out what inorder means.
- It means: leftNode - root - rightNode
- Use recursion

# Approach:
- Use a helper function recursiveTraversal that takes a node and a pointer to the result slice.
- Base case: If the node is nil, return.
- Recursively traverse the left subtree.
- Append the current node's value to the result slice.
- Recursively traverse the right subtree.
- The main function inorderTraversal initializes an empty result slice, calls the helper, and returns the result.

# Key:
- Recursion: The recursive approach naturally follows the inorder pattern by leveraging the call stack to track nodes. Each recursive call handles a subtree, ensuring the left-root-right order.
- Pointer to Slice: Passing a pointer to the result slice (*[]int) allows the helper function to modify the same slice across recursive calls, avoiding the need to return and merge slices.
- Base Case: Checking for nil nodes prevents infinite recursion and handles empty trees or missing children.
- Go Idioms: Using a pointer to the result slice is a common Go pattern for modifying slices in recursive functions, ensuring efficient memory usage.

# Complexity:
- Time: $O(n)$ where $n$ is the number of nodes.
- Space: $O(h)$ where $h$ is the height of the tree.

# Constraints:
- The number of nodes in the tree is in the range [0, 100].
- Node values are integers in the range [-100, 100].
- The input tree is valid.

# Optimizations:
- Iterative Solution: The follow-up suggests an iterative approach using an explicit stack to mimic the recursive call stack. This could reduce the risk of stack overflow for very deep trees (though unlikely given the constraint of up to 100 nodes). An iterative solution would have the same time and space complexity but might be preferred in some production environments.

- Memory Efficiency: The current solution is efficient, but pre-allocating the result slice with a known size (if the tree size is provided) could reduce reallocations during append.

- Morris Traversal: An advanced optimization (not implemented here) could achieve $$O(1)$$ auxiliary space by temporarily modifying the tree structure to avoid using a stack or recursion. This is complex and may not be practical for most use cases.

# List of solved problems:

- 1\. Two Sum
- 6\. Zigzag Conversion
- 9\. Palindrome Number
- 13\. Roman to Integer
- 14\. Longest Common Prefix
- 20\. Valid Parentheses
- 21\. Merge Two Sorted Lists
- 26\. Remove Duplicates from Sorted Array
- 27\. Remove Element
- 28\. Find the Index of the First Occurrence in a String
- 58\. Length of Last Word
- 66\. Plus One
- 67\. Add Binary
- 88\. Merge Sorted Array
- 94\. Binary Tree Inorder Traversal
- 108\. Convert Sorted Array to Binary Search Tree
- 202\. Happy Number
- 263\. Ugly Number
- 412\. Fizz Buzz
- 977\. Squares of a Sorted Array
