package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表的头节点 head ，判断链表中是否有环。
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head

	for slow != nil {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	// fast走在前面，判断fast就行了
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
