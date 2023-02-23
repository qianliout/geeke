package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		node = nil
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
