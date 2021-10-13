package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

func isSameTree(p *treenode.TreeNode, q *treenode.TreeNode) bool {
	// 这种解法就是bfs，深度递归
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 也可以使用bfs进行实现，但是golang标准库中没有实现双端队列（dqueue）,所以，要用bfs还要使用双端队列，
// 或者使用两个队列
