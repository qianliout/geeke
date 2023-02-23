package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
func recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, last := head, head.Next
	for last != nil && last.Next != nil {
		last = last.Next
		pre = pre.Next
	}
	pre.Next = nil
	next := recursion(head)
	last.Next = next
	return last
}
