// Package lengthOfLastWord contains a stub for LeetCode problem 0058. lengthOfLastWord.
//
// Link: https://leetcode.com/problems/lengthOfLastWord/
//
// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.
//
// Example 1:
//
// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:
//
// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:
//
// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

package lengthOfLastWord

func lengthOfLastWord(s string) int {
	length := 0
	// Walk backwards, skip trailing spaces
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length > 0 {
				break // finished counting last word
			}
			continue // skip trailing spaces
		}
		length++
	}
	return length
}
