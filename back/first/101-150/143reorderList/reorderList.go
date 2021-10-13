package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	root := &listnode.ListNode{Val: 1}
	root.Next = &listnode.ListNode{Val: 2}
	root.Next.Next = &listnode.ListNode{Val: 3}
	root.Next.Next.Next = &listnode.ListNode{Val: 4}
	root.Next.Next.Next.Next = &listnode.ListNode{Val: 5}
	reorderList(root)
	listnode.PrintListNode(root)
}

/*

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/

func reorderList(head *listnode.ListNode) {
	//fmt.Println("root value ", head.Val)
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	tailPre := head
	// 这里可以使用map优化
	for tailPre.Next.Next != nil {
		tailPre = tailPre.Next
	}

	// 这三步的顺序不可以变的,这里有个小技巧,1,置为空的,放在最后,2,先处理后面的,再处理前面的,从前到后
	tailPre.Next.Next = head.Next
	head.Next = tailPre.Next
	tailPre.Next=nil
	reorderList(head.Next.Next)
}
