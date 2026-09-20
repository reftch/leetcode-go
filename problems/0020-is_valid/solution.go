// Package is_valid contains a stub for LeetCode problem 0020. Valid Parentheses.
//
// Link: https://leetcode.com/problems/valid-parentheses/
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
// Output: true
//
// Example 2:
//
// Input: s = "()[]{}"
// Output: true
//
// Example 3:
//
// Input: s = "(]"
// Output: false
//
// Example 4:
//
// Input: s = "([])"
// Output: true
//
// Example 5:
//
// Input: s = "([)]"
// Output: false

package is_valid

// is_valid reports whether s is a valid parentheses string.
func is_valid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, r := range s {
		switch r {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != byte(r) {
				return false
			}
		default:
			return false
		}
	}
	return len(stack) == 0
}
