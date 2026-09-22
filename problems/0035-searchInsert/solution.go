// Package searchInsert contains a stub for LeetCode problem 0035. searchInsert.
//
// Link: https://leetcode.com/problems/searchInsert/
//
// Given a sorted array of distinct integers and a target value,
// return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Example 1:
//
// Input: nums = [1,3,5,6], target = 5
// Output: 2
// Example 2:
//
// Input: nums = [1,3,5,6], target = 2
// Output: 1
// Example 3:
//
// Input: nums = [1,3,5,6], target = 7
// Output: 4

package searchInsert

func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}
	return i
}
