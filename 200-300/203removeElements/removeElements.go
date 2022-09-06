package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {
	dump := new(ListNode)
	dump.Next = head
	cur := dump
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}

	return dump.Next
}
