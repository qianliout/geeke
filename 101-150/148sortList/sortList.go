package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := ListNode{Val: 4}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next = &ListNode{Val: 2}
	res := sortList(&head)
	PrintListNode(res)
}

//使用归并排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 先分隔
	slow, fast := head, head.Next
	// fast.Next!=nil的目的就是为了使fast保存的最后一个结点
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// 保存分隔结果:切开链表
	mid := slow.Next
	slow.Next = nil
	// 递归左右两边
	left, right := sortList(head), sortList(mid)

	// 归并
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

	// 把剩下的合并
	if left != nil {
		cur.Next = left
	} else {
		cur.Next = right
	}
	return dump.Next
}

//归并排序-迭代法
func sortList2(head *ListNode) *ListNode {
	return nil
}
