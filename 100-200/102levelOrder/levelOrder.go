package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
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
	return ans
}
