package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	return Divide(lists, 0, len(lists)-1)
}

func mergeKLists3(lists []*ListNode) *ListNode {
	dump := &ListNode{}
	for i := range lists {
		dump.Next = mergeTwoLists(dump, lists[i])
	}
	return dump.Next
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// 因为都是已排序链表，可以使用分治法,这是最快的方法
func Divide(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	leftNode := Divide(lists, left, mid)
	rightNode := Divide(lists, mid+1, right)
	// 以下这种写法不对,得想明白其中原因,
	// leftNode := Divide(lists, left, mid-1)
	// rightNode := Divide(lists, mid, right)

	// 这样写不对的原因就是：mid := left + (right-left)/2 这种方式取的是小值,而递归退出条件是left<=right
	// 比如有5个元素，0，1，2，3，4
	// 第一次循环【0，1】【2，3，4】
	// [0,1]和下一次循环时，mid=0，两组是[0,-1],[0,1]这样就进入了死循环了

	return mergeTwoLists(leftNode, rightNode)
}
