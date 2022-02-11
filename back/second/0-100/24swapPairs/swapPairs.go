package main

import (
	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {
	head := &listnode2.ListNode{Val: 3}
	head.Next = &listnode2.ListNode{Val: 5}
	head.Next.Next = &listnode2.ListNode{Val: 4}
}

func swapPairs(head *listnode2.ListNode) *listnode2.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(listnode2.ListNode)
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
