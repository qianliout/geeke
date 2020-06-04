package main

import (
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
	var p *ListNode
	for fast != nil && fast.Next != nil {
		p = slow
		slow = slow.Next
		fast = fast.Next.Next
		p.Next = pre
		pre = p
	}

	// slow表示的是后半段的开始, 如果总的节点是偶数个数,那么fast就没有到最后,slow就是在前半段的最后一个,所以要处理一下
	if fast != nil {
		slow = slow.Next
	}
	// 然后比较
	for p != nil && slow != nil {
		if p.Val != slow.Val {
			return false
		}
		p = p.Next
		slow = slow.Next
	}
	if p != nil || slow != nil {
		return false
	}
	return true
}

// 左侧指针
var left *ListNode

func isPalindrome2(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	// 后序遍历代码
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}
