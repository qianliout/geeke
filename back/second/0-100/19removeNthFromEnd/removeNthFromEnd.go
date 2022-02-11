package main

import (
	"fmt"

	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {
	head := &listnode2.ListNode{Val: 1}
	head.Next = &listnode2.ListNode{Val: 2}
	head.Next.Next = &listnode2.ListNode{Val: 3}
	head.Next.Next.Next = &listnode2.ListNode{Val: 4}
	head.Next.Next.Next.Next = &listnode2.ListNode{Val: 5}
	removeNthFromEnd(head, 2)
}

/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
*/
func removeNthFromEnd(head *listnode2.ListNode, n int) *listnode2.ListNode {
	if head == nil || n == 0 {
		return head
	}
	dump := new(listnode2.ListNode)
	dump.Next = head
	i, pre, cur := 0, dump, head
	for cur != nil {
		cur = cur.Next
		i++
		if i > n {
			pre = pre.Next
		}
	}
	fmt.Println(pre.Val)
	pre.Next = pre.Next.Next
	return dump.Next
}
