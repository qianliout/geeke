package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 1}
	head.Next = &listnode.ListNode{Val: 2}
	head.Next.Next = &listnode.ListNode{Val: 3}
	head.Next.Next.Next = &listnode.ListNode{Val: 4}
	head.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	res := rotateRight2(head, 2)
	listnode.PrintListNode(res)
}

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
https://leetcode-cn.com/problems/rotate-list
*/

func rotateRight(head *listnode.ListNode, k int) *listnode.ListNode {
	if head == nil {
		return nil
	}
	// 先计算有多少个节点{
	num := 0
	n := head
	for n != nil {
		num++
		n = n.Next
	}
	k = k % num
	if k == 0 {
		return head
	}
	interval := 0
	first := head
	second := head
	for second.Next != nil {
		if interval < k {
			second = second.Next
			interval++
		} else {
			first = first.Next
			second = second.Next
		}
	}
	ans := first.Next
	first.Next = nil
	second.Next = head

	return ans
}
