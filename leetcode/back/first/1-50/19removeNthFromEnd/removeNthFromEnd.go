package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	list1 := listnode.ListNode{Val: 1}
	list1.Next = &listnode.ListNode{Val: 2}
	list1.Next.Next = &listnode.ListNode{Val: 3}
	list1.Next.Next.Next = &listnode.ListNode{Val: 4}
	list1.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	res := removeNthFromEnd2(&list1, 2)
	listnode.PrintListNode(res)
}

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
*/
//方法一，两次遍历
func removeNthFromEnd(head *listnode.ListNode, n int) *listnode.ListNode {
	if head == nil {
		return head
	}
	if n == 0 {
		return head
	}
	first := head
	length := 0
	dump := new(listnode.ListNode)
	dump.Next = head
	for first != nil {
		length++
		first = first.Next
	}
	length -= n
	first = dump
	for length > 0 {
		first = first.Next
		length--
	}
	first.Next = first.Next.Next
	return dump.Next
}

// 双指针
func removeNthFromEnd2(head *listnode.ListNode, n int) *listnode.ListNode {
	if head == nil || n == 0 {
		return head
	}
	dump := new(listnode.ListNode) // 增加哑节点，避免了好多麻烦
	dump.Next = head
	first := dump
	interval := 0
	remove := dump
	for first.Next != nil {
		first = first.Next
		interval++
		if interval > n {
			remove = remove.Next
		}
	}
	remove.Next = remove.Next.Next
	return dump.Next
}
