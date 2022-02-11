package array

import (
	"qianliout/leetcode/common/array"
)

func ReverseLinkedList(head *array.ListNode) *array.ListNode {
	cur := head
	var prev *array.ListNode
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
