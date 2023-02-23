package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	flag := true
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
	for len(stark) > 0 {
		l := len(stark)
		th := make([]int, l)
		for i := 0; i < l; i++ {
			pop := stark[0]
			stark = stark[1:]
			if flag {
				th[i] = pop.Val
			} else {
				th[l-i-1] = pop.Val
			}
			if pop.Left != nil {
				stark = append(stark, pop.Left)
			}
			if pop.Right != nil {
				stark = append(stark, pop.Right)
			}
		}
		flag = !flag
		ans = append(ans, th)
	}
	return ans
}
