package tree

import "outback/leetcode/common/treenode"

/*
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/
func BuildTree105(preorder []int, inorder []int) *treenode.TreeNode {
	inpos := make(map[int]int, 0)
	for key, value := range inorder {
		inpos[value] = key
	}
	length := len(preorder) - 1
	return BuildTreeHelper105(preorder, inorder, inpos, 0, length, 0, length)
}
// not finished 没有完成呢
func BuildTreeHelper105(preorder, inorder []int, inps map[int]int, pStart, pEnd, IStart, Iend int) *treenode.TreeNode {
	if pStart == pEnd {
		return nil
	}
	root := treenode.TreeNode{Val: preorder[pStart]}
	rootIndex := inps[preorder[pStart]]
	leftLength := rootIndex - IStart

	root.Left = BuildTreeHelper105(preorder, inorder, inps, pStart+1, pStart+leftLength+1, IStart, rootIndex)
	root.Right = BuildTreeHelper105(preorder, inorder, inps, pStart+leftLength+1, pEnd, rootIndex+1, Iend)
	return &root
}
