/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
package main

func longestConsecutive(nums []int) int {
	parent := make(map[int]int, len(nums))
	size := make(map[int]int, len(nums))

	keys := make(map[int]bool, len(nums))
	for _, num := range nums {
		keys[num] = true
		parent[num] = num
		size[num] = 1
	}

	for key := range keys {
		if _, ok := keys[key+1]; ok {
			parent[key] = key + 1
			size[parent[key]] += size[key]
		}
	}

}

// @lc code=end
