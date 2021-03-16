package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	list1 := ListNode{Val: 2}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}
	res := reverseKGroup(&list1, 2)
	PrintListNode(res)
}

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
示例 :
给定这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5
说明 :
    你的算法只能使用常数的额外空间。
    你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
// 递归实现
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前k个元素
	newHead := reverse(a, b)
	// 递归反转后面的
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转以 a 为头结点的链表
func reverse(head, b *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	var nxt *ListNode

	cur, nxt = head, head
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

/*
用栈，我们把 k 个数压入栈中，然后弹出来的顺序就是翻转的！
这里要注意几个问题：
第一，剩下的链表个数够不够 k 个（因为不够 k 个不用翻转）；
第二，已经翻转的部分要与剩下链表连接起来。
*/
func reverseKGroupUseStack(head *ListNode, n int) *ListNode {
	res := new(ListNode)
	end := res.Next
	if head == nil {
		return head
	}
	stack := make([]int, 0)
	curr := head
	for curr != nil {
		// push to stack
		revers := true
		for i := 0; i < n; i++ {
			stack = append(stack, curr.Val)
			curr = curr.Next
		}
		// pup to stack
		for len(stack) == n {
			val := stack[len(stack)-1]
			res.Next = &ListNode{Val: val}
			stack = stack[:len(stack)-1]
		}
	}
	return res.Next
}

func reverse2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var pre *ListNode
	for curr != nil {
		tem := curr.Next
		curr.Next = pre
		pre = curr
		curr = tem
	}
	return curr
}
