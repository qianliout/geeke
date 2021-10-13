package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 3}
	head.Next = &listnode.ListNode{Val: 5}
	head.Next.Next = &listnode.ListNode{Val: 4}
}

func swapPairs(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(listnode.ListNode)
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
