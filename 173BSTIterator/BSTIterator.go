package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。
注意，指针初始化为一个不存在于 BST 中的数字，所以对 next() 的首次调用将返回 BST 中的最小元素。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 的中序遍历中至少存在一个下一个数字。
*/

type BSTIterator struct {
	stark []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{stark: make([]*TreeNode, 0)}
	// 把左子树全部入stark
	for root != nil {
		bst.stark = append(bst.stark, root)
		root = root.Left
	}
	return bst
}

func (this *BSTIterator) Next() int {
	last := this.stark[len(this.stark)-1]
	this.stark = this.stark[:len(this.stark)-1]
	node := last.Right
	// 把当前的root及其左子树全部入stark
	for node != nil {
		this.stark = append(this.stark, node)
		node = node.Left
	}
	return last.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stark) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
