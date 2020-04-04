package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 1}
	//root.Right.Left = &TreeNode{Val: 1}
	//root.Left.Left = &TreeNode{Val: 1}
	//root.Left.Right = &TreeNode{Val: 3}
	res := rob2(root)
	fmt.Println("res is ", res)
}

/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
示例 1:
输入: [3,2,3,null,3,null,1]
     3
    / \
   2   3
    \   \
     3   1
输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
示例 2:
输入: [3,4,5,1,3,null,1]
     3
    / \
   4   5
  / \   \
 1   3   1
输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
*/

//典型的dp
var mermery map[*TreeNode][]int

// todo 这种方法不能得到正确的答案，是为什么呢
func rob(root *TreeNode) int {
	//if mermery == nil {
	//	mermery = make(map[*TreeNode]int)
	//}
	if root == nil {
		return 0
	}
	//if v, ok := mermery[root]; ok {
	//	return v
	//}

	next := 0
	if root.Left != nil {
		next = rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		next = rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 抢的情况
	doit := root.Val + next
	// 不抢的情况
	notdoit := rob(root.Left) + rob(root.Right)
	this := Max(doit, notdoit)

	//mermery[root] = this
	return this
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, nus := range nums {
		if nus > max {
			max = nus
		}
	}
	return max

}

func rob2(root *TreeNode) int {
	// 动态规划
	mermery = make(map[*TreeNode][]int)
	non, rob := robTree(root)

	return Max(non, rob)
}
func robTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	if v, ok := mermery[root]; ok {
		return v[0], v[1]
	}
	// 左子树
	ln, lr := robTree(root.Left) // ln表示不抢，lr表示抢
	// 右子树
	rn, rr := robTree(root.Right)

	// 不抢, 抢
	l := Max(ln, lr) + Max(rn, rr)
	r := root.Val + ln + rn
	mermery[root] = []int{l, r}

	return Max(ln, lr) + Max(rn, rr), root.Val + ln + rn
}
