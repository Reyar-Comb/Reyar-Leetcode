/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */
package main

// @lc code=start
func subarraySum(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	ans := 0
	hash := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		if val, ok := hash[nums[i]+k]; ok {
			ans += val
		}
		hash[nums[i]]++
	}
	if val, ok := hash[k]; ok {
		ans += val
	}
	return ans
}

// @lc code=end
