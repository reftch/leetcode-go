// Package longestCommonPrefix contains a stub for LeetCode problem 0014. LongestCommonPrefix.
//
// Link: https://leetcode.com/problems/longestCommonPrefix/
package longestCommonPrefix

import "fmt"

// longestCommonPrefix implements the solution.
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the common prefix
	prefix := strs[0]

	// Compare with remaining strings
	for i := 1; i < len(strs); i++ {
		for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			fmt.Printf("DEBUG:   - Mismatch found! Shrinking prefix from %q to %q\n",
				prefix, prefix[:len(prefix)-1])
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}
