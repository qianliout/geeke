package main

import (
	"expvar"

	"golang.org/x/crypto/curve25519"

	. "outback/leetcode/common/listnode"
)

func main() {

}

/*
请判断一个链表是否为回文链表。
示例 1:
输入: 1->2
输出: false
示例 2:
输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

// 把数据存在一个slice里,然后使用双指针的方式进行判断,这样做很简单
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 方法二,使用快慢指针翻转前半部分,然后和后半部分比较
func recursion(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		// 翻转前半部门
		tem := slow.Next
		slow.Next = pre
		pre = slow
		slow = tem
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 然后比较
	first := head

}
