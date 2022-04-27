package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {

}

func reorderList(head *ListNode) {
	// dfs(head)
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	pre, last := head, head.Next
	for last != nil && last.Next != nil {
		pre = pre.Next
		last = last.Next
	}
	pre.Next = nil
	next := head.Next

	head.Next = last
	last.Next = next
	reorderList(next)
}

func dfs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	pre, last := head, head.Next
	for last != nil && last.Next != nil {
		pre = pre.Next
		last = last.Next
	}
	next := head.Next
	pre.Next = nil
	head.Next = last
	last.Next = dfs(next)
	return head
}
