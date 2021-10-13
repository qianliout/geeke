package dfsbfs

import (
	treenode2 "outback/leetcode/back/common/treenode"
)

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

func FindIndex(list []int, value int) int {
	for key, v := range list {
		if v == value {
			return key
		}
	}
	return -1
}

// 由前序遍历preorder的第一个节点即为根节点
// 因为树中没有重复元素，所以由中序遍历以根节点左右两边分别是左右子树
// 结合前序遍历，知道左右子树分别的前序遍历和中序遍历，递归调用即可

// 在前序遍历知道根节点之后，要在中序遍历中找到根节点，避免每次都做线性查找
// 可以使用哈希表来存储中序遍历中每个节点对应的下标
// 使用辅助哈希表和递归过程的栈空间，所以Time:(n),n为节点数量 Space:O(n)
func BuildTreePreAndIn(preorder []int, inorder []int) *treenode2.TreeNode {
	inPos := make(map[int]int) // 辅助hash表存储中序遍历节点的下标
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i // 中序遍历元素作为key下标作为value存储
	}
	return BuildTreePreAndInHelper(preorder, 0, len(preorder)-1, 0, inPos)
}
func BuildTreePreAndInHelper(pre []int, preStart, preEnd, inStart int, inPos map[int]int) *treenode2.TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := treenode2.TreeNode{Val: pre[preStart]}
	rootIndex := inPos[pre[preStart]]
	leftLen := rootIndex - inStart
	root.Left = BuildTreePreAndInHelper(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = BuildTreePreAndInHelper(pre, preStart+leftLen+1, preEnd, rootIndex+1, inPos)
	return &root
}

// pre前序遍历序列，preStart开始下标，preEnd结束下标，中序遍历序列inStart开始下标，和hash表
func buildTreeHelper(pre []int, preStart, preEnd, inStart int, inPos map[int]int) *treenode2.TreeNode {
	if preStart > preEnd { // 如果前序遍历序列的开始下标已经大于结束下标
		return nil // 说明本次构建的二叉树中已经没有节点，返回空树
	}
	// 否则拿出前序遍历序列开始的元素，构建一个树节点，即为当前树的根节点
	root := &treenode2.TreeNode{Val: pre[preStart]}
	// 在hash表中查到当前根节点在中序遍历中的下标
	rootIdx := inPos[pre[preStart]]
	// 减去本次中序遍历的开始下标，即为左子树的节点个数
	leftLen := rootIdx - inStart
	// 使用节点个数在前序遍历中确定左右子树的边界
	// 递归的调用自己分别构建左右子树
	root.Left = buildTreeHelper(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildTreeHelper(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root // 最后返回这棵树即可
}

func BuildTree(preorder, inorder []int) *treenode2.TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历第一个值为根节点
	root := treenode2.TreeNode{Val: preorder[0]}
	// 因为没有重复元素，所以可以直接根据值来查找根节点在中序遍历中的位置
	mid := FindIndex(inorder, preorder[0])
	root.Left = BuildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])
	return &root
}
