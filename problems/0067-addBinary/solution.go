// Package addBinary contains a stub for LeetCode problem 0067. addBinary.
//
// Link: https://leetcode.com/problems/addBinary/
//
// Given two binary strings a and b, return their sum as a binary string.
// Example 1:
//
// Input: a = "11", b = "1"
// Output: "100"
// Example 2:
//
// Input: a = "1010", b = "1011"
// Output: "10101"
//
// Constraints:
//
// 1 <= a.length, b.length <= 104
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the zero itself.

package addBinary

import "fmt"

func addBinary(a string, b string) string {
	var getVal func(s string, i int) int
	getVal = func(s string, i int) int {
		len_s := len(s)
		if i < len_s && s[len_s-i-1] == 49 {
			return 1
		}
		return 0
	}

	a3 := 0
	sum := ""
	for i := range max(len(a), len(b)) {
		a1 := getVal(a, i)
		a2 := getVal(b, i)

		val := a1 + a2 + a3

		if val > 2 {
			sum = "1" + sum
			a3 = 1
		} else if val > 1 {
			sum = "0" + sum
			a3 = 1
		} else {
			sum = fmt.Sprintf("%d%s", val, sum)
			a3 = 0
		}
	}

	if a3 == 1 {
		sum = "1" + sum
	}

	return sum
}
