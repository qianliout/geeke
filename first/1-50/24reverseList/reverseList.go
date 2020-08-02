package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	res := reverseList3(head)
	PrintListNode(res)
}

/*
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

// 第一种方法，使用栈，先把链表的数据存入栈，再把数取出来重新构造链表。简单
func reverseList(head *ListNode) *ListNode {
	return head
}

// 方法二，一次遍历
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var curr *ListNode
	pre := head
	for pre != nil {
		t := pre.Next
		pre.Next = curr
		curr = pre
		pre = t
	}
	return curr
}

//这个递归解法一定要理解
//https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/bu-bu-chai-jie-ru-he-di-gui-di-fan-zhuan-lian-biao/
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
