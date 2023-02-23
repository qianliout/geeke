package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
	if root == nil {
		return ans
	}
	for len(stark) > 0 {
		l := len(stark)
		th := make([]int, 0)
		for i := 0; i < l; i++ {
			pop := stark[0]
			stark = stark[1:]
			th = append(th, pop.Val)
			if pop.Left != nil {
				stark = append(stark, pop.Left)
			}
			if pop.Right != nil {
				stark = append(stark, pop.Right)
			}
		}
		ans = append(ans, th)
	}
	res := make([][]int, len(ans))
	for i := range ans {
		res[len(ans)-i-1] = ans[i]
	}
	return res
}
