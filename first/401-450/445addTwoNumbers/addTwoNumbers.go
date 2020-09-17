package main

import (
	"fmt"
	
	. "outback/leetcode/common/listnode"
)

func main() {
	
}

/*
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。
进阶：
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。
示例：
输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s := UseSlice(l1, l2)
	fmt.Println("s is ", s)
	
	return makeNode(UseSlice(l1, l2))
}

// use slice ,stark最好
func UseSlice(l1, l2 *ListNode) []int {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	cur := l1
	for cur != nil {
		list1 = append(list1, cur.Val)
		cur = cur.Next
	}
	// fmt.Println("list1 is ", list1)
	
	cur = l2
	
	for cur != nil {
		list2 = append(list2, cur.Val)
		cur = cur.Next
	}
	// fmt.Println("list2 is ", list2)
	pre := 0
	if len(list2) > len(list1) {
		list2, list1 = list1, list2
	}
	res := make([]int, len(list1)+1)
	
	i := 0
	for i < len(list2) {
		v := pre + list2[len(list2)-i-1] + list1[len(list1)-i-1]
		if v >= 10 {
			pre = v / 10
			res[len(res)-i-1] = v % 10
		} else {
			pre = 0
			res[len(res)-i-1] = v
		}
		i++
	}
	for i < len(list1) {
		v := pre + list1[len(list1)-i-1]
		if v >= 10 {
			pre = v / 10
			res[len(res)-i-1] = v % 10
		} else {
			pre = 0
			res[len(res)-i-1] = v
		}
		i++
	}
	if pre > 0 {
		res[0] = pre
		return res
	}
	return res[1:]
}

func makeNode(nums []int) *ListNode {
	dum := new(ListNode)
	if len(nums) == 0 {
		return dum.Next
	}
	
	cur := dum
	for _, v := range nums {
		no := &ListNode{Val: v}
		cur.Next = no
		cur = cur.Next
	}
	return dum.Next
}

// 另外一种方法，反转两个链表，然后想加


