# leet-code-in-go
A series of leet code problems solutions using Golang. Not in any particular order.
Each branch of this repo is responsible for 1 problem, and represents a solution for this problem.
On each branch, different commits may exist with different implementations.

# Problem: 20\. Valid Parentheses
- Given a string s containing just the characters '\(', '\)', '\{', '\}', '\[' and '\]', determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

# First thoughts:
- Analyze parts of the string, checking if the order of open/closed brackets is correct.
- Keep track of where the current order starts and ends.
- Use a hashMap to count every open bracket.
- Two pointers.
- I could not think of an efficient logic, my previous attempt generated too many tests conditions, so I had to search the answer.
- Using LIFO, of course...
- Use a stack to keep track of open brackets, use a hashMap to close brackets and remove them from stack.

# Approach:
- Initialize a stack (slice) to track opening brackets.
- Use a map (map[rune]rune) to store matching pairs.
- Iterate through each character in the string:
- - If it's an opening bracket (\(, \{, \[), push its corresponding closing bracket onto the stack.
- - If it's a closing bracket (\), \}, \]), check if it matches the top of the stack:
- - Mismatch? → Invalid (return false).
- - Match? → Pop from the stack.
- Final check: If the stack is empty, all brackets were properly matched (return true).

# Key:
- Use a stack (LIFO).
- Hash maps for fast lookup.

# Complexity:
- Time: $$O(n)$$
- Space: $$O(n)$$

# Constraints:
- (1 <= s.length <= 104)
- s consists of parentheses only '\(\)\[\]\{\}'.

# Optimizations:
- Early exit: If the string length is odd → Immediately false.
- Fixed capacity: Pre-allocate stack space (make([]rune, 0, len(s)/2)) for efficiency.

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
- 66\. Plus One
- 67\. Add Binary
- 202\. Happy Number
- 263\. Ugly Number
- 412\. Fizz Buzz
