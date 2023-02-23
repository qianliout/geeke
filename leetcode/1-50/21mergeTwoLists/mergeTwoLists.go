package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := new(ListNode)
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dump.Next
}

// 使用递归呢
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dump := new(ListNode)
	if list1.Val <= list2.Val {
		dump.Next = &ListNode{Val: list1.Val}
		list1 = list1.Next
		dump.Next.Next = mergeTwoLists2(list1, list2)
		return dump.Next
	} else {
		dump.Next = &ListNode{Val: list2.Val}
		list2 = list2.Next
		dump.Next.Next = mergeTwoLists2(list1, list2)
		return dump.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
