package main

import (
	treenode2 "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
调用 next() 将返回二叉搜索树中的下一个最小的数。
示例：
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
链接：https://leetcode-cn.com/problems/binary-search-tree-iterator
*/
type BSTIterator struct {
	root  *treenode2.TreeNode
	stack []int
}

func Constructor(root *treenode2.TreeNode) BSTIterator {
	stack := make([]int, 0)
	inorder(root, &stack)
	// sort.Ints(stack)
	return BSTIterator{root: root, stack: stack}
}

// 递归解法.因为是二叉搜索树迭代器，所以使用中序遍历最好
func inorder(root *treenode2.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		inorder(root.Right, res)
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if len(this.stack) == 0 {
		return 0
	}
	v := this.stack[0]
	this.stack = this.stack[1:len(this.stack)]
	return v
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
