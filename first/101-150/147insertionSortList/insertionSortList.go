package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := ListNode{Val: 4}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next = &ListNode{Val: 2}
	res := insertionSortList(&head)
	PrintListNode(res)
}

/*
插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。
插入排序算法：
    插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
    每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
    重复直到所有输入数据插入完为止。
链接：https://leetcode-cn.com/problems/insertion-sort-list
*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode) // 维护一个已排序的一个List
	for head != nil {
		curr := head
		head = head.Next
		p := search(dump, curr)
		// 插入当前curr结点
		curr.Next = p.Next
		p.Next = curr
	}
	return dump.Next
}

// 返回插入节点插入位置的前一个节点
func search(head, node *ListNode) *ListNode {
	// 为什么使用head.Next去比较呢，因为head有一个dump结点
	for head != nil && head.Next != nil && head.Next.Val < node.Val {
		head = head.Next
	}
	return head
}
