package main

import (
	listnode2 "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/
func reverseKGroup(head *listnode2.ListNode, k int) *listnode2.ListNode {
	start, nexStart := head, head
	for i := 0; i < k; i++ {
		if nexStart == nil {
			return head
		}
		nexStart = nexStart.Next
	}

	hd := reverse(start, nexStart)
	start.Next = reverseKGroup(nexStart, k)
	return hd
}

// 反转链表
func reverse(head, tail *listnode2.ListNode) *listnode2.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *listnode2.ListNode
	cur := head
	for cur != tail {
		nex := cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
