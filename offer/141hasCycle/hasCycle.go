package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

func hasCycle(head *ListNode) bool {
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

func detectCycle(head *ListNode) *ListNode {
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

var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return recursion(head)
}

// 使用全局变量
func recursion(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := recursion(right.Next)
	res = res && right.Val == left.Val
	left = left.Next
	return res
}
