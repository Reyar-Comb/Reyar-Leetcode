/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package main

// @lc code=start
func trap(height []int) int {
	leftMax := 0
	left := 0
	rightMax := 0
	right := len(height) - 1
	sum := 0
	for {
		if left >= right {
			break
		}
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}

		if leftMax < rightMax {
			sum += leftMax - height[left]
			left++
		} else {
			sum += rightMax - height[right]
			right--
		}
	}
	return sum
}

// @lc code=end
