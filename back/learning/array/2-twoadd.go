package array

import (
	"qianliout/leetcode/back/common/listnode"
)

func addTwoNumbers1(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	result := new(listnode.ListNode)
	var cur = result

	for l1 != nil || l2 != nil {
		var a, b int = 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		val := a + b
		if cur.Next == nil {
			cur.Next = &listnode.ListNode{}
		}
		if val+cur.Next.Val >= 10 {
			val -= 10
			cur.Next.Next = &listnode.ListNode{Val: 1}
		}
		cur.Next.Val += val
		cur = cur.Next
	}
	// 这一步是思考的关键
	return cur.Next
}

func addTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	result := new(listnode.ListNode)
	cur := result

	for l1 != nil || l2 != nil {
		var a, b int = 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		if cur.Next == nil {
			cur.Next = &listnode.ListNode{}
		}
		if a+b+cur.Next.Val >= 10 {
			cur.Next.Next.Val = 1
			cur.Next.Val = a + b + cur.Next.Val - 10
		}

	}

}
