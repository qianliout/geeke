package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

后序遍历 postorder = [9,15,7,20,3]
中序遍历 inorder = [9,3,15,20,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := find(inorder, postorder[len(postorder)-1])
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return root
}

func find(inorde []int, val int) int {
	for i, v := range inorde {
		if v == val {
			return i
		}
	}
	return -1
}
