# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 88\. Merge Sorted Array
- You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
- Merge nums1 and nums2 into a single array sorted in non-decreasing order.
- The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.  

- Example 1:

Input: `nums1 = [1,2,3,0,0,0]`, `m = 3`, `nums2 = [2,5,6]`, `n = 3`  
Output: `[1,2,2,3,5,6]`  
Explanation: The arrays we are merging are `[1,2,3]` and `[2,5,6]`.  
The result of the merge is `[1,2,2,3,5,6]` with the underlined elements coming from `nums1`.  

- Example 2:  

Input: `nums1 = [1]`, `m = 1`, `nums2 = []`, `n = 0`  
Output: `[1]`  
Explanation: The arrays we are merging are `[1]` and `[]`.  
The result of the merge is `[1]`.  

- Example 3:  

Input: `nums1 = [0]`, `m = 0`, `nums2 = [1]`, `n = 1`  
Output: `[1]`  
Explanation: The arrays we are merging are [] and [1].  
The result of the merge is [1].  
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.  

# First thoughts:
- The problem requires merging two sorted arrays into `nums1` without returning a new array, suggesting an in-place solution to save space.
- Since `nums1` has extra space (filled with zeros), we can use it to store the merged result, but merging from the start risks overwriting `nums1`’s elements before they’re processed.
- A naive approach would use an extra array to merge and copy back, but this uses O(m + n) space, which we want to avoid.
- Working backwards from the end of `nums1` seems promising, as it uses the empty space (zeros) first and avoids overwriting unprocessed elements.

# Approach:
- Use three pointers: `p1` for the last valid element in `nums1` (index `m-1`), `p2` for the last element in `nums2` (index `n-1`), and `p` for the last position in `nums1` (index `m+n-1`).
- Compare `nums1[p1]` and `nums2[p2]`, placing the larger element at `nums1[p]` and decrementing the corresponding pointer (`p1` or `p2`) and `p`.
- Continue until one array is exhausted (`p1 < 0` or `p2 < 0`).
- If `p2 >= 0`, copy remaining `nums2` elements to `nums1` starting at `p`.
- No need to copy remaining `nums1` elements, as they’re already in place.
- Handle edge cases like `n = 0` (no changes needed) or `m = 0` (copy `nums2` to `nums1`).

# Key:
- **In-Place Merging**: Work backwards to use `nums1`’s extra space without overwriting unprocessed elements.
- **Pointer Management**: Use `p1`, `p2`, and `p` to track positions in `nums1`, `nums2`, and the merged result, respectively.
- **Edge Cases**: Handle empty arrays (`m = 0` or `n = 0`) to ensure robustness.
- **Sorted Property**: Leverage the fact that both arrays are sorted to compare only the largest remaining elements.

# Complexity:
- Time: $$O(m + n)$$
- Space: $$O(1)$$

# Constraints:
- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`

# Optimizations:
- **Early Exit**: If `n = 0`, return immediately since `nums1` is already correct.
- **No Extra Array**: By merging backwards, we avoid the need for an auxiliary array, achieving \( O(1) \) space.
- **Minimize Comparisons**: Compare only the largest remaining elements, leveraging the sorted nature of the arrays.
- **No Unnecessary Copies**: Skip copying `nums1`’s remaining elements, as they’re already in their final positions.
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
