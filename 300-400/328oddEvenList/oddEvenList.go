package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {
	list := &ListNode{Val: 2}
	list.Next = &ListNode{Val: 3}
	list.Next.Next = &ListNode{Val: 4}
	after := oddEvenList(list)
	PrintListNode(after)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, event := head, head.Next
	eventStart := event
	for event != nil && event.Next != nil {
		odd.Next = event.Next
		odd = odd.Next
		event.Next = odd.Next
		event = event.Next
	}

	odd.Next = eventStart
	return head
}
