package main

import (
	"fmt"

	. "outback/leetcode/common/listnode"
)

func main() {

}

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	nums1 := make([]int, 0)
	nums2 := make([]int, 0)

	cur := l1
	for cur != nil {
		nums1 = append(nums1, cur.Val)
		cur = cur.Next
	}
	cur = l2
	for cur != nil {
		nums2 = append(nums2, cur.Val)
		cur = cur.Next
	}
	dump := new(ListNode)
	pre := 0
	curNode := dump
	for len(nums1) > 0 || len(nums2) > 0 {
		r := pre
		if len(nums1) > 0 {
			r += nums1[0]
			nums1 = nums1[1:]
		}
		if len(nums2) > 0 {
			r += nums2[0]
			nums2 = nums2[1:]
		}
		if r >= 10 {
			r = r - 10
			pre = 1
		} else {
			pre = 0
		}
		node := ListNode{Val: r}
		curNode.Next = &node
		curNode = curNode.Next
	}
	if pre != 0 {
		curNode.Next = &ListNode{Val: pre}
	}

	return dump.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dump := new(ListNode)
	cur := dump
	pre := 0
	for l1 != nil || l2 != nil {
		a := pre
		if l1 != nil {
			a += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			a += l2.Val
			l2 = l2.Next
		}
		if a >= 10 {
			a = a - 10
			pre = 1
		} else {
			pre = 0
		}
		cur.Next = &ListNode{Val: a}
		cur = cur.Next
	}
	if pre != 0 {
		cur.Next = &ListNode{Val: pre}
	}
	return dump.Next
}
