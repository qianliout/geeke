package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	ans := reverseKGroup(head, 2)
	PrintListNode(ans)
}

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换
*/
func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}
	cru, n := head, 1
	for cru != nil && cru.Next != nil && n%k > 0 {
		cru = cru.Next
		n++
	}
	if n < k {
		return head
	}

	nexNode := cru.Next

	cru.Next = nil

	nex := reverseKGroup2(nexNode, k)

	res := reverseUseRecursion(head)

	now := res
	for now != nil && now.Next != nil {
		now = now.Next
	}
	now.Next = nex
	return res
}

func reverseUseRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nex := head.Next
	head.Next = nil
	last := reverseUseRecursion(nex)
	cur := last
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head // 不会为nil
	return last
}

// 这种方法更优
func reverseKGroup(head *ListNode, k int) *ListNode {
	start, nextStart := head, head
	for i := 0; i < k; i++ {
		if nextStart == nil {
			return head
		}
		nextStart = nextStart.Next
	}
	ans := reverse(start, nextStart)
	start.Next = reverseKGroup(nextStart, k)
	return ans
}

// 反转链表
func reverse(head, tail *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != tail {
		nex := cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
