package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	InOrderTraversal(root)
}

/*
根据一棵树的中序遍历与后序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/
// 树中没有重复的元素
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	indexMap := make(map[int]int)
	for k, v := range inorder {
		indexMap[v] = k
	}
	return build(&postorder, 0, len(inorder)-1, indexMap)
}

// 这道题的关键就是postorder一定要传指针，因为left递归和right递归都要去改这个值，不然就会有重复值
func build(postorder *[]int, leftIndex, rightIndex int, indexMap map[int]int) *TreeNode {
	if leftIndex > rightIndex {
		return nil
	}
	val := (*postorder)[len(*postorder)-1]
	*postorder = (*postorder)[:len(*postorder)-1]
	root := &TreeNode{Val: val}
	index := indexMap[val]
	// 注意，一定是root.Right在前，因为是通过 val := (*postorder)[len(*postorder)-1]这一句确定root的值，所以一定要先确定right,
	// 然后把postorder的right部分都删除完成之后才能确定left
	root.Right = build(postorder, index+1, rightIndex, indexMap)
	root.Left = build(postorder, leftIndex, index-1, indexMap)
	return root
}

// 这才是这道题的正解，也是最好理解的答案
func BuildTree(inorder []int, postorder []int) *TreeNode {

	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历第一个值为根节点
	root := TreeNode{Val: postorder[len(postorder)-1]}
	// 因为没有重复元素，所以可以直接根据值来查找根节点在中序遍历中的位置
	mid := FindIndex(inorder, postorder[len(postorder)-1])
	r := len(inorder) - mid - 1 // 右边元素的个数，这个数字一定要好好理解
	root.Right = BuildTree(inorder[mid+1:], postorder[len(postorder)-1-r:len(postorder)-1])
	root.Left = BuildTree(inorder[:mid], postorder[:len(postorder)-1-r])
	return &root
}

func FindIndex(list []int, value int) int {
	for key, v := range list {
		if v == value {
			return key
		}
	}
	return -1
}
