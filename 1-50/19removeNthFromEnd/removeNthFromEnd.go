package main

import (
	"fmt"

	. "qianliout/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	end := removeNthFromEnd(head, 4)
	PrintListNode(end)

}

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	all := 0
	for cur != nil {
		all++
		cur = cur.Next
	}
	fmt.Println("all is ", all)

	dump := new(ListNode)
	dump.Next = head
	preNode := dump
	pre := all - n

	for pre > 0 && preNode != nil {
		preNode = preNode.Next
		pre--
	}

	if preNode != nil && preNode.Next != nil {
		preNode.Next = preNode.Next.Next
	}
	return dump.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if n == 0 {
		return head
	}
	first := head
	length := 0
	dump := new(ListNode)
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
