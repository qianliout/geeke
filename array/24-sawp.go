package array

import "outback/leetcode/common/array"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func SwapLinkedList(head *array.ListNode) *array.ListNode {
	list := &array.ListNode{Next: head}
	for prev, node := list, list.Next; node != nil; node = node.Next {
		if node.Next != nil {
			//prev.Next, treenode.Next, treenode.Next.Next = treenode.Next, treenode.Next.Next, treenode
			swapNode(prev, node, node.Next)
			prev = node
		}
	}
	return list.Next
}

func swapNode(prev, node, next *array.ListNode) {
	prev.Next = next
	node.Next = next.Next
	next.Next = node
}
