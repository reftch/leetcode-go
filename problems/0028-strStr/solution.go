// Package strStr contains a stub for LeetCode problem 0028. strStr.
//
// Link: https://leetcode.com/problems/strStr/
//
// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
//
// Example 1:
//
// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0
// Explanation: "sad" occurs at index 0 and 6.
// The first occurrence is at index 0, so we return 0.
// Example 2:

// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
// Explanation: "leeto" did not occur in "leetcode", so we return -1.

package strStr

func strStr(haystack string, needle string) int {
	haystack_len := len(haystack)
	needle_len := len(needle)

	if needle_len == 0 {
		return 0
	}

	if haystack_len < needle_len {
		return -1
	}

	for i := 0; i <= haystack_len-needle_len; i++ {
		mathes := true
		for j := range needle_len {
			if haystack[i+j] != needle[j] {
				mathes = false
				break
			}
		}
		if mathes {
			return i
		}
	}
	return -1
}
