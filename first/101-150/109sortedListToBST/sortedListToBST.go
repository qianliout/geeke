package main

import (
	. "outback/leetcode/common/listnode"
	. "outback/leetcode/common/treenode"
)

func main() {
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}
	res := sortedListToBST(head)
	InOrderTraversal(res)
	//r := InorderTraversalToSlice(res)
	//fmt.Println(r)
}

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
给定的有序链表： [-10, -3, 0, 5, 9],
一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：
      0
     / \
   -3   9
   /   /
 -10  5
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
*/

// 通过map构造
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	length := -1
	cur := head
	listMap := make(map[int]*ListNode)
	for cur != nil {
		length++
		listMap[length] = cur
		cur = cur.Next
	}
	root := buildTree(listMap, 0, length)
	return root
}

func buildTree(listMap map[int]*ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	root := new(TreeNode)
	mid := start + (end-start)/2
	// 这里是这道题的关键点，-10，-3，两个数，我们应该选后面一个作为root点，但是start + (end-start)/2这种方式会选前面一个
	if (end-start)&1 == 1 {
		mid = start + (end+1-start)/2
	}
	if n, ok := listMap[mid]; ok {
		root.Val = n.Val
	}
	root.Left = buildTree(listMap, start, mid-1)
	root.Right = buildTree(listMap, mid+1, end)
	return root
}
