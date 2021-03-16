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
	dump.Next = head
	pre := dump

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// 交换
		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = second
		head = first.Next
	}

	return dump.Next
}

func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nex := head.Next
	last := head.Next.Next
	nex.Next = head
	nex.Next.Next = swapPairs3(last)
	return nex
}

func swapPairs4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode)
	pre := dump
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		pre.Next = second
		second.Next = first
		first.Next = second.Next

		pre = second
		head = first.Next
	}

	return dump.Next
}
