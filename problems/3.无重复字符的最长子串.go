/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	index := make([]int, 128)
	for i := range index {
		index[i] = -1
	}
	left := 0
	right := 0
	max := 0
	for right <= len(s)-1 {
		if index[s[right]] >= left {
			left = index[s[right]] + 1
		}
		index[s[right]] = right
		right++
		if right-left > max {
			max = right - left
		}
	}
	return max
}

// @lc code=end
