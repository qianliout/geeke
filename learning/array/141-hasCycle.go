package array

import (
	"fmt"
	"outback/leetcode/common/array"
)

/*
给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
*/
func HasCycle(head *array.ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil && fast.Next != nil {
		if slow == fast {
			fmt.Println(slow.Val, fast.Val)
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 使用set的方法
func HasCycleBySet(head *array.ListNode) bool {
	isExst := make(map[int]interface{}, 0)

	for head != nil {
		if _, ok := isExst[head.Val]; ok {
			return true
		} else {
			isExst[head.Val] = nil
			head = head.Next
		}
	}
	return false
}
