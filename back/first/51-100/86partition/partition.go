package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 1}
	head.Next = &listnode.ListNode{Val: 4}
	head.Next.Next = &listnode.ListNode{Val: 3}
	head.Next.Next.Next = &listnode.ListNode{Val: 2}
	head.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &listnode.ListNode{Val: 2}

	res := partition(head, 3)
	listnode.PrintListNode(res)

}

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
示例:
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
*/
func partition(head *listnode.ListNode, x int) *listnode.ListNode {
	if head == nil {
		return head
	}

	dump1 := new(listnode.ListNode)
	dump2 := new(listnode.ListNode)
	last1 := dump1
	last2 := dump2

	curr := head

	for curr != nil {
		if curr.Val < x {
			last1.Next = curr
			last1 = last1.Next
		} else {
			last2.Next = curr
			last2 = last2.Next
		}
		curr = curr.Next
	}
	last1.Next = dump2.Next
	last2.Next = nil
	return dump1.Next
}
