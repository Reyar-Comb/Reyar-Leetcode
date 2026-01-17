package main

import (
	"fmt"
)

func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}}
	addTwoNumbers(&l1, &l2)
	node := l1
	for {
		fmt.Printf("%d->", node.Val)
		if node.Next == nil {
			break
		}
		node = *node.Next
	}
	fmt.Println("")
	node = l2
	var a []int
	for {
		a = append(a, node.Val)
		if node.Next == nil {
			break
		}
		node = *node.Next
	}
	fmt.Println(a)
}
