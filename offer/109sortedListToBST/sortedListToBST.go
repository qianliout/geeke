package main

import (
	. "outback/leetcode/common/listnode"
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
      0
     / \
   -3   9
   /   /
 -10  5
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return new(TreeNode)
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	dump := new(ListNode)
	dump.Next = head

	pre, slow, fast := dump, head, head
	for fast != nil && fast.Next != nil {
		pre, slow, fast = pre.Next, slow.Next, fast.Next.Next
	}
	pre.Next = nil
	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(slow.Next),
	}
	return root
}
