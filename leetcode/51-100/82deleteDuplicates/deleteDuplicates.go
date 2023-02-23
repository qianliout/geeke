package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表
*/
// 使用递归才是正确的解法
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates(head.Next)
	} else {
		move := head.Next
		for move != nil && move.Val == head.Val {
			move = move.Next
		}
		return deleteDuplicates(move)
	}
	return head
}
