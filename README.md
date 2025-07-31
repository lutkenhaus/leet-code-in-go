# leet-code-in-go
- A series of leet code problems solutions using Golang. Not in any particular order.
- Each branch is related to one problem, and represents a solution for this problem.
- Different commits may exist with different implementations.
- Branches can contain code explanations in the `README.md` file.

# Problem: 6\. Zigzag Conversion
- The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:  
```
P   A   H   N
A P L S I I G
Y   I   R
```
- And then read line by line: "PAHNAPLSIIGYIR"
- Write the code that will take a string and make this conversion given a number of rows:  
```
string convert(string s, int numRows);
```

## Example 1:

Input: s = "PAYPALISHIRING", numRows = 3  
Output: "PAHNAPLSIIGYIR"  

## Example 2:

Input: s = "PAYPALISHIRING", numRows = 4  
Output: "PINALSIGYAHRPI"  
Explanation:  
```
P     I    N
A   L S  I G
Y A   H R
P     I
```

## Example 3:

Input: s = "A", numRows = 1  
Output: "A"  

# First thoughts:
-

# Approach:
-

# Key:
-

# Complexity:
- Time: $$O(1 + n)$$
- Space: $$O(1 + n)$$

# Constraints:
- (1 <= s.length <= 1000)
- s consists of English letters (lower-case and upper-case), ',' and '.'.
- (1 <= numRows <= 1000)

# Optimizations:
-

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
