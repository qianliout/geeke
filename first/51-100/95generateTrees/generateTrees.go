package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
示例:
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return helper(0, n)
}

func helper(start, end int) []*TreeNode {
	// n 表示总数，i 表示第i 做根结点
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	if start == end {
		node := TreeNode{Val: start}
		res = append(res, &node)
		return res
	}

	for i := start; i <= end; i++ {
		leftNodes := helper(start, i-1)
		rightNodes := helper(i+1, end)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				node := &TreeNode{Val: i}
				node.Left = leftNode
				node.Right = rightNode
				res = append(res, node)
			}
		}
	}
	return res
}
