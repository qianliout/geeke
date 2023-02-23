package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 1}
	head.Next = &listnode.ListNode{Val: 2}
	head.Next.Next = &listnode.ListNode{Val: 3}
	head.Next.Next.Next = &listnode.ListNode{Val: 3}
	head.Next.Next.Next.Next = &listnode.ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	res := deleteDuplicates2(head)
	listnode.PrintListNode(res)
}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:
输入: 1->1->1->2->3
输出: 2->3
*/
// 这种写法不符合题义，因为只是删除了重复的，但还是保留了一个
func deleteDuplicates(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return head
	}
	dump := new(listnode.ListNode)
	dump.Next = head
	lastNode := dump.Next

	curr := head.Next
	for curr != nil {
		if curr.Val != lastNode.Val {
			lastNode.Next = curr
			lastNode = curr
		}
		curr = curr.Next
	}
	return dump.Next
}

func deleteDuplicates2(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return head
	}
	dump := new(listnode.ListNode)
	dump.Next = head
	cur := dump
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tem := cur.Next
			for tem != nil && tem.Next != nil && tem.Val == tem.Next.Val {
				tem = tem.Next
			}
			cur.Next = tem.Next
		} else {
			cur = cur.Next
		}
	}
	return dump.Next
}
