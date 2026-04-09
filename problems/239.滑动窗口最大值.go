/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
package main

// @lc code=start
type Mqueue struct {
	data []int
}

func (q *Mqueue) push(num int) {
	for len(q.data) > 0 && num > q.data[len(q.data)-1] {
		q.pop_back()
	}
	q.data = append(q.data, num)
}

func (q *Mqueue) pop_back() {
	q.data = q.data[:len(q.data)-1]
}

func (q *Mqueue) pop_head() {
	q.data = q.data[1:]
}

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := 0
	mq := Mqueue{}
	ans := []int{}
	for ; right < k; right++ {
		mq.push(nums[right])
	}
	ans = append(ans, mq.data[0])
	for {
		if right >= len(nums) {
			break
		}
		if left >= 0 && nums[left] == mq.data[0] {
			mq.pop_head()
		}
		mq.push(nums[right])
		ans = append(ans, mq.data[0])

		left++
		right++
	}
	return ans
}

// @lc code=end
