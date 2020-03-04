package listnode

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
