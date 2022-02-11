package main

import (
	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {

}

func hasCycle(head *listnode2.ListNode) bool {
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

func detectCycle(head *listnode2.ListNode) *listnode2.ListNode {
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

var left *listnode2.ListNode

func isPalindrome(head *listnode2.ListNode) bool {
	left = head
	return recursion(head)
}

// 使用全局变量
func recursion(right *listnode2.ListNode) bool {
	if right == nil {
		return true
	}
	res := recursion(right.Next)
	res = res && right.Val == left.Val
	left = left.Next
	return res
}
