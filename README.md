# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 58\. Length of Last Word
- Given a string s consisting of words and spaces, return the length of the last word in the string.
- A word is a maximal substring consisting of non-space characters only.

# First thoughts:
- Analyze the string from the end. (inverted loop)
- Stop when space is found.
- Trim extra spaces.
- Count the letters.

# Approach:
- Trim Leading/Trailing Spaces:
- - Use strings.TrimSpace() to eliminate unnecessary spaces at the start and end of the string.
- Backward Iteration:
- - Loop from the end of the trimmed string until a space is encountered.
- - Count the length of the last word during this iteration.

# Key:
- Efficient Traversal: By processing the string from the end, we stop early once the last word is counted.
- Space Handling: TrimSpace ensures we ignore irrelevant leading/trailing spaces.

# Complexity:
- Time: $$O(n)$$
- Space: $$O(1)$$

# Constraints:
- (1 <= s.length <= 10^4)
- s consists of only English letters and spaces ' '.
- There will be at least one word in s.

# Optimizations:
- Don't use strings package, instead loop twice.
- First loop skips the spaces at the end.
- Second loop counts the word.

# List of solved problems:

- 1\. Two Sum
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
- 202\. Happy Number
- 263\. Ugly Number
- 412\. Fizz Buzz
