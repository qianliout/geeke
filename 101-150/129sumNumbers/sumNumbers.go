package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	//root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}
	//root.Right.Right = &TreeNode{Val: 4}
	ans := sumNumbers(root)
	fmt.Println("ans is ", ans)
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := make([][]int, 0)
	used := make(map[*TreeNode]bool)
	path := make([]int, 0)
	dfs(root, used, path, &res)

	ans := 0
	//fmt.Println("res ia ", res)
	for _, nums := range res {
		a := 0
		for _, nun := range nums {
			a = a*10 + nun
		}
		ans = ans + a
	}

	return ans
}

func dfs(root *TreeNode, used map[*TreeNode]bool, path []int, res *[][]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, append(append([]int{}, path...), root.Val))
		return
	}
	if !used[root] {
		used[root] = true
		path = append(path, root.Val)
		dfs(root.Left, used, path, res)
		dfs(root.Right, used, path, res)
		used[root] = true
		path = path[:len(path)-1]
	}
}
