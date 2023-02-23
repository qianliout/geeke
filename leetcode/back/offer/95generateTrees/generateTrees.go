package main

import (
	treenode2 "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树 。
示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateTrees(n int) []*treenode2.TreeNode {
	if n == 0 {
		return make([]*treenode2.TreeNode, 0)
	}
	return helper(1, n)
}

// func helper(start, end int) []*TreeNode {
// 	// n 表示总数，i 表示第i 做根结点
// 	res := make([]*TreeNode, 0)
// 	// 这个基本条件是怎么判断的呢，为什么要加一个nil呢，要好好理解
// 	// 结合下面这一名rightNodes := helper(i+1, end),也就是当每次的i=n时，只是左边有数据，左边没有子树，那么就应该是一个nil,如果不写这个，那么也就不会循环左边的循环了，
// 	if start > end {
// 		res = append(res, nil)
// 		fmt.Println("start end ", start, end, len(res))
// 		return res
// 	}
// 	for i := start; i <= end; i++ {
// 		leftNodes := helper(start, i-1)
// 		rightNodes := helper(i+1, end)
// 		for _, leftNode := range leftNodes {
// 			for _, rightNode := range rightNodes {
// 				node := &TreeNode{Val: i}
// 				node.Left = leftNode
// 				node.Right = rightNode
// 				res = append(res, node)
// 			}
// 		}
// 	}
// 	return res
// }

func helper(start, end int) []*treenode2.TreeNode {
	res := make([]*treenode2.TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftNodes := helper(start, i-1)
		rightNodes := helper(i+1, end)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				res = append(res, &treenode2.TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return res
}

func numTrees(n int) int {
	dp := make([]int, n+3)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ { // right
		for j := 1; j <= i; j++ { // left
			// 这里要好好理解，意思是选j这个值做为node结点的值，那么left就有dp[j-1],right就有dp[i-j]个
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
