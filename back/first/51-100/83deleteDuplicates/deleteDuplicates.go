package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 1}
	head.Next = &listnode.ListNode{Val: 1}
	head.Next.Next = &listnode.ListNode{Val: 1}
	res := deleteDuplicates(head)
	listnode.PrintListNode(res)

}

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
示例 1:
输入: 1->1->2
输出: 1->2
示例 2:
输入: 1->1->2->3->3
输出: 1->2->3
*/
// 排序链表，所以就是操作排序的能力
// 先使用map
func deleteDuplicates(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	exit := make(map[int]bool)
	cur := head.Next

	if cur != nil && cur.Next != nil {
		if exit[cur.Val] {
			cur.Next = cur.Next.Next
		} else {
			exit[cur.Val] = true
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicates2(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head

	// fixme 以下这种写法，为什么就不对呢
	//pre, cur := head, head.Next
	//for cur != nil {
	//	if cur.Val != pre.Val {
	//		pre.Next = cur
	//		pre = pre.Next
	//	}
	//	cur = cur.Next
	//
	//}
	//return pre
}
