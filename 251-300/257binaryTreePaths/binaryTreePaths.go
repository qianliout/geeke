package main

import (
	"strconv"

	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。
示例:
输入:
   1
 /   \
2     3
 \
  5
输出: ["1->2->5", "1->3"]
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
来源：力扣（LeetCode）
*/
func binaryTreePaths(root *TreeNode) []string {
	res := make([][]int, 0)
	if root == nil {
		return []string{}
	}
	used := make(map[*TreeNode]bool)
	dfs(root, used, []int{}, &res)
	stringRes := make([]string, 0)
	for _, nums := range res {
		s := ""
		for _, num := range nums {
			if len(s) == 0 {
				s = strconv.Itoa(num)
			} else {
				s += "->" + strconv.Itoa(num)
			}
		}
		stringRes = append(stringRes, s)
	}

	return stringRes
}

func dfs(root *TreeNode, used map[*TreeNode]bool, path []int, res *[][]int) {
	if root == nil {
		//*res = append(*res, append([]int{}, path...))
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
		used[root] = false
		path = path[:len(path)-1]
	}
}
