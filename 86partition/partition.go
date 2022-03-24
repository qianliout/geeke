package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置
*/
func partition(head *ListNode, x int) *ListNode {
	cur, dump1, dump2 := head, new(ListNode), new(ListNode)
	cur1, cur2 := dump1, dump2
	for cur != nil {
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur1.Next
		} else {
			cur2.Next = cur
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	cur1.Next = dump2.Next
	cur2.Next = nil
	return dump1.Next
}
