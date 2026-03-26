/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
package main

func longestconsecutive(nums []int) int {
	keys := make(map[int]bool, len(nums))
	for _, num := range nums {
		keys[num] = true
	}
	max := 0
	for key := range keys {
		if _, ok := keys[key-1]; ok {
			continue
		}
		length := 1
		for keys[key+length] {
			length++
		}
		if length > max {
			max = length
		}
	}
	return max
}

// @lc code=end
