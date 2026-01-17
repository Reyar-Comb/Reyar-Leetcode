/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	var c int
	var prec int
	for {
		c = (node1.Val + node2.Val + prec) / 10
		node2.Val = (node1.Val + node2.Val + prec) % 10
		prec = c
		if node1.Next == nil && node2.Next == nil {
			if c == 0 {
				return l2
			}
			node2.Next = &ListNode{c, nil}
			return l2
		}
		if node1.Next == nil {
			node2 = node2.Next
			for {
				c = (node2.Val + prec) / 10
				node2.Val = (node2.Val + prec) % 10
				prec = c
				if node2.Next == nil && prec != 1 {
					return l2
				} else if node2.Next == nil && prec == 1 {
					node2.Next = &ListNode{1, nil}
					return l2
				}
				node2 = node2.Next
			}
		}
		if node2.Next == nil {
			node2.Next = node1.Next
			node2 = node2.Next
			for {
				c = (node2.Val + prec) / 10
				node2.Val = (node2.Val + prec) % 10
				prec = c
				if node2.Next == nil && prec != 1 {
					return l2
				} else if node2.Next == nil && prec == 1 {
					node2.Next = &ListNode{1, nil}
					return l2
				}
				node2 = node2.Next
			}
		}

		node1 = node1.Next
		node2 = node2.Next

	}
}

// @lc code=end
