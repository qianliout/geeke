package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	list1 := ListNode{Val: 2}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}
	list1.Next.Next.Next = &ListNode{Val: 6}
	list1.Next.Next.Next.Next = &ListNode{Val: 7}
	res := swapPairs(&list1)
	PrintListNode(res)
}

/*
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/

func swapPairs(head *ListNode) *ListNode {
	return Iteration(head)
}

// 递归解法
func swap(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode := head
	secondeNode := head.Next
	firstNode.Next = swap(secondeNode.Next)
	secondeNode.Next = firstNode
	return secondeNode
}

// 迭代解法
func Iteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	prevNode := dummy
	for head != nil && head.Next != nil {

		firstNode := head
		secondNode := head.Next

		// swap
		prevNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prevNode = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}

// 这种写法好难理解啊
func SwapLinkedList(head *ListNode) *ListNode {
	list := &ListNode{Next: head}
	for prev, node := list, list.Next; node != nil; node = node.Next {
		if node.Next != nil {
			// prev.Next, treenode.Next, treenode.Next.Next = treenode.Next, treenode.Next.Next, treenode
			swapNode(prev, node, node.Next)
			prev = node
		}
	}
	return list.Next
}

func swapNode(prev, node, next *ListNode) {
	prev.Next = next
	node.Next = next.Next
	next.Next = node
}

func Iteration2(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for head != nil && head.Next != nil {
		first := head
		nxt := head.Next
		third := head.Next.Next

		pre.Next = nxt
		first.Next = third
		nxt.Next = first

		pre = head
		head = first.Next
	}
	return dummy.Next

	/*
		dummy := new(ListNode)
		dummy.Next = head
		prevNode := dummy
		for head != nil && head.Next != nil {

			firstNode := head
			secondNode := head.Next

			// swap
			prevNode.Next = secondNode
			firstNode.Next = secondNode.Next
			secondNode.Next = firstNode

			prevNode = firstNode
			head = firstNode.Next
		}
		return dummy.Next
	*/
}

// 这种解法是错的，但是没有找到错的原因
func swapPairs3(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for head != nil && head.Next != nil {
		nxt := head.Next
		third := head.Next.Next

		pre.Next = nxt
		nxt.Next = head
		head.Next = nxt.Next

		head = third
		pre = head
	}
	return dummy.Next
}
