/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := -1
	for {
		if left > right {
			break
		}
		if (right-left)*min(height[left], height[right]) > max {
			max = (right - left) * min(height[left], height[right])
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

// @lc code=end
