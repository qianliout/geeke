package main

import (
	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {

}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5            :=-c
示例 2:
输入: 1->1->1->2->3
输出: 2->3
*/
// 删除重复的，
func deleteDuplicates(head *listnode2.ListNode) *listnode2.ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 把重复的删除
func deleteDuplicates2(head *listnode2.ListNode) *listnode2.ListNode {
	dump := new(listnode2.ListNode)
	dump.Next = head
	cur := dump
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tem := cur.Next
			for tem != nil && tem.Next != nil && tem.Next.Val == tem.Val {
				tem = tem.Next
			}
			cur.Next = tem.Next
		} else {
			cur = cur.Next
		}
	}
	return dump.Next
}
