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
- My first instinct was to use a matrix, but then I realized I had no knowledge of the number of collumns, only rows.
- So I decided to use a hash map, storing each row as a key, and characters as rune values.
- Since there are only 2 directions in the zigzag logic (up/down), a boolean flag is sufficient.

# Approach:
- Edge Handling:  
If numRows == 1 or numRows >= len(s), return s immediately (no zigzag needed).  
- Row Storage:  
Use a slice of strings.Builder (optimized for dynamic string building in Go).  
- Direction Logic:  
Traverse the string while flipping direction (+1 → down, -1 → up) at the first/last row.  
- Concatenation:  
Merge all rows into the final result.

# Key:
- Analyze each character in the string sequentially.
- Implement logic to determine whether the current direction is up or down (using a toggle or counter).

# Complexity:
- Time: $$O(n)$$
- Space: $$O(n)$$

# Constraints:
- (1 <= s.length <= 1000)
- s consists of English letters (lower-case and upper-case), ',' and '.'.
- (1 <= numRows <= 1000)

# Optimizations:
- Why use a hash map if an array can do the job? strings.Builder is slightly more efficient in Go.
- Using an integer instead of a bool can eliminate one if condition.

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
- 202\. Happy Number
- 263\. Ugly Number
- 412\. Fizz Buzz
