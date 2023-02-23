package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dump := new(ListNode)
	cur := dump

	for l1 != nil || l2 != nil {
		v := 0
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if cur.Next == nil {
			cur.Next = &ListNode{}
		}
		if v+cur.Next.Val >= 10 {
			v -= 10
			cur.Next.Next = &ListNode{Val: 1}
		}
		cur.Next.Val += v
		cur = cur.Next
	}

	return dump.Next
}
