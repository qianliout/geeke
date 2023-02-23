package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur, sum := head, 0
	for cur != nil {
		sum++
		cur = cur.Next
	}

	// 优化，不然会超时
	for i := 0; i < k%sum; i++ {
		head = rotate(head)
	}
	return head
}

func rotate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode)
	dump.Next = head
	pre, last := dump, head
	for last != nil && last.Next != nil {
		pre = pre.Next
		last = last.Next
	}
	pre.Next = nil
	last.Next = head
	return last
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur, sum := head, 0
	for cur != nil {
		sum++
		cur = cur.Next
	}
	dump := new(ListNode)
	dump.Next = head
	pre, last := dump, head
	if k%sum == 0 {
		return head
	}

	// 优化，不然会超时
	for i := 0; i < sum-k%sum; i++ {
		pre, last = pre.Next, last.Next
	}
	pre.Next = nil

	bef := last
	for bef != nil && bef.Next != nil {
		bef = bef.Next
	}
	bef.Next = head
	return last
}
