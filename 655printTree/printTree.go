package main

import (
	"fmt"
	"strconv"
	"strings"

	. "qianliout/leetcode/common/treenode"
	. "qianliout/leetcode/common/utils"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	tree := printTree(root)
	for i := range tree {
		fmt.Println(strings.Join(tree[i], " \"\" "))
	}
}

func printTree(root *TreeNode) [][]string {
	height := deep(root)
	res := make([][]string, height)
	width := 1<<(height) - 1

	for i := range res {
		res[i] = make([]string, width)
	}
	dsf(root, &res, 0, (width-1)/2, height-1)
	return res
}

func dsf(root *TreeNode, ans *[][]string, r, c, height int) {
	if root == nil {
		return
	}
	(*ans)[r][c] = strconv.Itoa(root.Val)

	if root.Left != nil {
		dsf(root.Left, ans, r+1, c-(1<<(height-r-1)), height)
	}
	if root.Right != nil {
		dsf(root.Right, ans, r+1, c+(1<<(height-r-1)), height)
	}
}

func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := deep(root.Left) + 1
	right := deep(root.Right) + 1
	return Max(left, right)
}
