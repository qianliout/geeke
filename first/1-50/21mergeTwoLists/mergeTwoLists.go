package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

/*
21. 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
// 这种题典型的递归实现
//如果 l1 或者 l2 一开始就是 null ，那么没有任何操作需要合并，所以我们只需要返回非空链表。
//否则，我们要判断 l1 和 l2 哪一个的头元素更小，然后递归地决定下一个添加到结果里的值。
//如果两个链表都是空的，那么过程终止，所以递归过程最终一定会终止。

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		if l1.Val <= l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoLists(l2.Next, l1)
			return l2
		}
	}
}
