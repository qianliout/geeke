package main

import (
	"math"

	treenode2 "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
二叉树的根是数组中的最大元素。
左子树是通过数组中最大值左边部分构造出的最大二叉树。
右子树是通过数组中最大值右边部分构造出的最大二叉树。
通过给定的数组构建最大二叉树，并且输出这个树的根节点。
示例 ：
输入：[3,2,1,6,0,5]
输出：返回下面这棵树的根节点：
      6
    /   \
   3     5
    \    /
     2  0
       \
        1
*/

func constructMaximumBinaryTree(nums []int) *treenode2.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := findMax(nums)
	root := &treenode2.TreeNode{Val: nums[max]}

	root.Left = constructMaximumBinaryTree(nums[:max])
	root.Right = constructMaximumBinaryTree(nums[max+1:])
	return root
}
func findMax(nums []int) int {
	max := math.MinInt64
	mi := 0
	for i, v := range nums {
		if v > max {
			max = v
			mi = i
		}
	}
	return mi
}
