package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	second, third := head.Next, head.Next.Next

	second.Next = head
	head.Next = swapPairs2(third)
	return second
}

// 迭代法
func swapPairs(head *ListNode) *ListNode {
	dump := &ListNode{Next: head}
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	dumpCur, HeadCur := dump, head

	for HeadCur != nil && HeadCur.Next != nil {
		first, second, third := HeadCur, HeadCur.Next, HeadCur.Next.Next
		first.Next = third
		second.Next = first

		dumpCur.Next = second

		dumpCur = dumpCur.Next.Next
		HeadCur = third
	}
	return dump.Next
}

func swap() *ListNode {

}
