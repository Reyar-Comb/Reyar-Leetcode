/*
 * @lc app=leetcode.cn id=1929 lang=golang
 *
 * [1929] 数组串联
 */

// @lc code=start
func getConcatenation(nums []int) []int {
	nums2 := make([]int, 0, len(nums)*2)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)
	return nums2
}

// @lc code=end

