package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head.Next
	third := head.Next.Next
	cur.Next = head
	head.Next = swapPairs(third)
	return cur
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode)

	for head != nil && head.Next != nil {




	}

	return cur

}
