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
- Initialize a result array of size `n` to store the sorted squares.
- Use two pointers: `firstP` at the start (index 0) and `secondP` at the end (index `n-1`) of `nums`.
- Use a third pointer `resultP` starting at the end of the result array (index `n-1`).
- While `firstP <= secondP`:
  - Compare the squares of `nums[firstP]` and `nums[secondP]`.
  - Place the larger square at `result[resultP]`, increment `firstP` or decrement `secondP`, and decrement `resultP`.
- Return the result array.

# Key:
- **Two Pointers**: Use pointers at the start and end of the sorted input array to compare squares, leveraging the non-decreasing order.
- **Backward Filling**: Build the result array from the end to place larger squares first, ensuring sorted order without extra sorting.
- **Handle Negatives**: Account for negative numbers, as their squares can be larger (e.g., `(-4)^2 = 16 > 3^2 = 9`).
- **Edge Cases**: Single-element arrays and arrays with duplicates or negative numbers are handled by the two-pointer logic.

# Complexity:
- Time: $$O(n)$$
- Space: $$O(1)$$

# Constraints:
- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` is sorted in non-decreasing order.

# Optimizations:
- **Avoid Sorting**: Use two pointers to achieve \( O(n) \) time instead of squaring and sorting (\( O(n \log n) \)).
- **Single Pass**: Process each element once by comparing squares and placing them directly in sorted order.
- **No In-Place**: In-place modification is avoided since it would require re-sorting, which is less efficient.
- **Minimal Extra Space**: Only use the required \( O(n) \) space for the output array, with no additional data structures.

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
