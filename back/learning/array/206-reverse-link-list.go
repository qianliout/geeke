package array

import (
	array2 "qianliout/leetcode/back/common/array"
)

func ReverseLinkedList(head *array2.ListNode) *array2.ListNode {
	cur := head
	var prev *array2.ListNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
