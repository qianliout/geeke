package main

import (
	"fmt"

	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	res := sortList(head)
	fmt.Println(res.Val)
}

/*
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/
// 主要是常数空间,归并排数的思想
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 切开链表
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	left, right := sortList(head), sortList(mid)

	// 开始归并
	dump := new(ListNode)
	cur := dump

	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}
	// 没有归并完的弄完，这里直接一个if判断就可以，当然也可以使用for,结果是一样的
	if left != nil {
		cur.Next = left
		// left = left.Next
	}
	if right != nil {
		cur.Next = right
		// right = right.Next
	}
	return dump.Next
}
