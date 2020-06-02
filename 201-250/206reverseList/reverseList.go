package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

/*
反转一个单链表。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return recursion(head)
}

func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := recursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

func Iteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	var pre *ListNode
	for curr != nil {
		tem := curr.Next
		curr.Next = pre
		pre = curr
		curr = tem
	}
	return pre
}
