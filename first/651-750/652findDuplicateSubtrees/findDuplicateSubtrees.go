package main

import (
	"fmt"
	"strconv"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 0}
	root.Right.Right.Right = &TreeNode{Val: 0}
	nodes := findDuplicateSubtrees(root)
	result := make([]*TreeNode, 0)
	exit := make(map[string]int)
	for _, node := range nodes {
		s := serialize(node, &result, &exit)
		fmt.Println(s)
	}

}

/*
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
两棵树重复是指它们具有相同的结构以及相同的结点值。
示例 1：
        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
下面是两个重复的子树：
      2
     /
    4
和
    4
因此，你需要以列表的形式返回上述重复子树的根结点。
*/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {

	result := make([]*TreeNode, 0)
	exit := make(map[string]int)
	if root == nil {
		return result
	}
	serialize(root, &result, &exit)
	return result
}

func serialize(root *TreeNode, result *[]*TreeNode, exit *map[string]int) string {
	if root == nil {
		return "null"
	}
	// 这里一定要使用前序遍历,使用其他就会出错,因为前后会加Null (这才是这道题的易错点)
	s := strconv.Itoa(root.Val) + "," + serialize(root.Left, result, exit) + "," + serialize(root.Right, result, exit)

	(*exit)[s] += 1

	if (*exit)[s] == 2 {
		*result = append(*result, root)
	}
	return s
}


