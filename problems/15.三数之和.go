/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package main

import "slices"

// @lc code=start
func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}

		}
	}
	return ans
}

// @lc code=end
