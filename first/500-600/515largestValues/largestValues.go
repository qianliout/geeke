package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}

	res := largestValues(root)
	fmt.Println("res is ", res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
您需要在二叉树的每一行中找到最大的值。
示例：
输入:
          1
         / \
        3   2
       / \   \
      5   3   9
输出: [1, 3, 9]
*/

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	ans := level(root)
	fmt.Println("level is ", ans)

	for _, a := range ans {
		res = append(res, Max(a...))
	}
	return res
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func dfs(root *TreeNode, res *[]int) int {
	if root == nil {
		return math.MinInt64
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, root.Val)
		return root.Val
	}
	max := math.MinInt64
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	if left > right {
		max = left
	} else {
		max = right
	}

	// *res = append(*res, max)
	return max
}

// 方法一，层序遍历，然后取值就可以了
func level(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {

		this := make([]int, 0)
		thisNode := make([]*TreeNode, 0)
		for _, last := range queue {
			this = append(this, last.Val)
			if last.Left != nil {
				thisNode = append(thisNode, last.Left)
			}
			if last.Right != nil {
				thisNode = append(thisNode, last.Right)
			}
		}

		ans = append(ans, this)
		queue = thisNode

	}
	return ans
}
