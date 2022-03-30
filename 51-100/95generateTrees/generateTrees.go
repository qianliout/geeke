package main

import (
	"fmt"

	. "qianliout/leetcode/common/treenode"
)

func main() {
	nodes := generateTrees(3)
	fmt.Println("nodes is ", len(nodes))
}

/*
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
*/
func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}
func generate(left, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if left > right {
		// 这一步非常关键
		res = append(res, nil)
		return res
	}
	for i := left; i <= right; i++ {
		leftNodes := generate(left, i-1)
		rightNodes := generate(i+1, right)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				node := &TreeNode{Val: i, Left: leftNode, Right: rightNode}
				res = append(res, node)
			}
		}
	}
	return res
}
