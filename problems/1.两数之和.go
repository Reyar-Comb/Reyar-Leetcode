/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		_, ok := m[num]
		if ok {
			return []int{i, m[num]}
		} else {
			m[target-num] = i
		}
	}
	return nil
}

// @lc code=end
