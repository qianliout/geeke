package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {

}

func hasCycle(head *listnode.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *listnode.ListNode) *listnode.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// 用数组就太简单了

var left *listnode.ListNode

func isPalindrome(head *listnode.ListNode) bool {
	left = head
	return recursion(head)
}

// 使用全局变量
func recursion(right *listnode.ListNode) bool {
	if right == nil {
		return true
	}
	res := recursion(right.Next)
	res = res && right.Val == left.Val
	left = left.Next
	return res
}
