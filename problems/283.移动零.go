/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package main

// @lc code=start
func moveZeroes(nums []int) {
	left := 0
	right := 0
	for ; left < len(nums); left++ {
		if nums[left] == 0 {
			if right < left {
				right = left
			}
			for ; right < len(nums); right++ {
				if nums[right] != 0 {
					nums[left], nums[right] = nums[right], nums[left]
					break
				}
			}
		}
	}
}

// @lc code=end
