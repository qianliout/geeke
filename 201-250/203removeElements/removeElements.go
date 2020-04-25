package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	res := removeElements(head, 1)
	PrintListNode(res)
}

/*
删除链表中等于给定值 val 的所有节点。
示例:
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dump := new(ListNode)
	dump.Next = head
	curr := dump
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dump.Next
}
