# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 977\. Squares of a Sorted Array
- Given an integer array `nums` sorted in non-decreasing order, return an array of **the squares of each number** sorted in non-decreasing order.

- Example 1:  

Input: `nums = [-4,-1,0,3,10]`  
Output: `[0,1,9,16,100]`  
Explanation: After squaring, the array becomes `[16,1,0,9,100]`.  
After sorting, it becomes `[0,1,9,16,100]`.  

- Example 2:  

Input: `nums = [-7,-3,2,3,11]`  
Output: `[4,9,9,49,121]`  

# First thoughts:
- Create a new array.
- Iterate through `nums` and make calculations, storing the results in the new array.
- If I use two pointers I can get O(1) space complexity.
- First pointer points at the beginning of `nums`, second pointer at the end.
- Check for the greater square. 

# Approach:
-

# Key:
-

# Complexity:
- Time: $$O(n)$$
- Space: $$O(1)$$

# Constraints:
- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in non-decreasing order.

# Optimizations:
-

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
