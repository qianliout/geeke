package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	node := sortList(head)
	PrintListNode(node)
}

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
func sortList(head *ListNode) *ListNode {
	return merger(head)
}

// 这种方式可以得到结果，但是会超时
func dfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := dfs(head.Next)
	dump := new(ListNode)
	dump.Next = next
	pre, last := dump, dump.Next
	for last != nil && head.Val > last.Val {
		pre = pre.Next
		last = last.Next
	}
	pre.Next = head
	head.Next = last
	return dump.Next
}

func merger(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 这道题目的关键点就是fast = head.Next,如果fast=head会导致死循环
	// 因为如果只有两个节点之后，就没有办法分隔
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	mid := slow.Next
	slow.Next = nil
	left, right := merger(head), merger(mid)
	// 开始归并
	dump := new(ListNode)
	cur := dump
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dump.Next
}
