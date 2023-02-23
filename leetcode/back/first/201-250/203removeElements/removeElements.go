package main

import (
	"fmt"
	"time"

	"outback/leetcode/back/common/listnode"
)

func main() {
	head := &listnode.ListNode{Val: 1}
	head.Next = &listnode.ListNode{Val: 2}
	removeElements2(head, 1)
	//PrintListNode(res)
}

/*
删除链表中等于给定值 val 的所有节点。
示例:
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

func removeElements(head *listnode.ListNode, val int) *listnode.ListNode {
	if head == nil {
		return head
	}
	dump := new(listnode.ListNode)
	dump.Next = head
	curr := dump
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dump.Next
}

func removeElements2(head *listnode.ListNode, val int) *listnode.ListNode {
	if head == nil {
		return head
	}
	pre := new(listnode.ListNode)
	cur := head
	cout := 0
	// 这里怎么就不对了,怎么就形成环了,要好好理解
	for cur != nil {
		fmt.Println("hello ", pre.Val, "word ", cur.Val, "count", cout)
		time.Sleep(time.Second)
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			cur = cur.Next
		}
		pre = pre.Next
		//cur = cur.Next
		cout++
		if cout > 10 {
			break
		}
	}
	return pre.Next
}
