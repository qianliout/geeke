package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 4}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode)
	dump.Next = head

	pre := dump
	for head != nil && head.Next != nil {
		first := head
		second := head.Next
		// 交换
		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = second
		head = first.Next

	}

	return dump.Next
}
