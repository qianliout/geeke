package main

import (
	. "qianliout/leetcode/common/listnode"
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
给定一个单链表的头节点  head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
*/
func sortedListToBST1(head *ListNode) *TreeNode {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	pre := new(ListNode)
	pre.Next = head
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next
	pre.Next = nil // 断开链表

	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(right),
	}
	return root
}
