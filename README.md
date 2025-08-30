# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 108\. Convert Sorted Array to Binary Search Tree
- Given an integer array `nums` where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
- Height-Balanced:  
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.  

- Example 1:  
Input: nums = [-10,-3,0,5,9]  
Output: [0,-3,9,-10,null,5]  
Explanation: [0,-10,5,null,-3,null,9] is also accepted.  

- Example 2:  
Input: nums = [1,3]  
Output: [3,1]  
Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.  

# First thoughts:
- Array is sorted.
- Search for the binary tree algorithm and explanation.
- Study and memorize.
- Try to create a non recursive implementation.

# Approach:
- Recursive function: start from the middle of the array, create the left and right subtrees.
- Divide and conquer.


# Key:
- By consistently choosing the middle element of any given array as the root, we ensure that the tree remains height-balanced.  
The left and right subtrees will have roughly the same number of nodes, preventing the tree from becoming a linked list and satisfying the height balance condition.

# Complexity:
- Time: $$O(n)$$
- Space: $$O(log n)$$

# Constraints:
- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in a strictly increasing order.

# Optimizations:
- Slicing Overhead: In Go, slicing (nums[:middle]) creates a new slice header which is very efficient, but still involves a small overhead. For maximum performance in very large arrays, the algorithm could be refactored to use start and end indices on the original array to avoid creating numerous slice headers. This would be a memory optimization, but the time complexity remains the same.

- Non-Recursive Implementation: As noted in your first thoughts, an iterative solution using a stack or queue to simulate the recursion is possible. However, the recursive solution is very intuitive, easy to understand, and naturally fits the divide-and-conquer paradigm for this problem. The logarithmic stack depth makes the recursive approach efficient and acceptable.

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
- 202\. Happy Number
- 263\. Ugly Number
- 412\. Fizz Buzz
- 977\. Squares of a Sorted Array
