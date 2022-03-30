package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {

}

/*
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cru := head
	for cru != nil && cru.Val == head.Val {
		cru = cru.Next
	}
	head.Next = deleteDuplicates(cru)
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	last, cur := head, head

	for cur != nil {
		for cur != nil && cur.Val == last.Val {
			cur = cur.Next
		}
		last.Next = cur
		last = last.Next
	}
	return head
}
