package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {

}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil
	next = reverse(next)
	for head != nil && next != nil {
		if head.Val != next.Val {
			return false
		}
		head, next = head.Next, next.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverse(head.Next)

	head.Next.Next = head
	head.Next = nil
	return next
}
