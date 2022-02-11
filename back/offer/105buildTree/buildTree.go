package main

import (
	treenode2 "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

func buildTree(preorder []int, inorder []int) *treenode2.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &treenode2.TreeNode{Val: preorder[0]}
	idx := findindex(inorder, preorder[0])

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return root
}

// 没有重复元素
func findindex(inorder []int, val int) int {
	for i, v := range inorder {
		if v == val {
			return i
		}
	}
	return 0
}
