package main

import (
	"fmt"
	"strconv"

	treenode2 "qianliout/leetcode/common/treenode"
)

func main() {
	root := &treenode2.TreeNode{Val: 1}
	root.Left = &treenode2.TreeNode{Val: 2}
	root.Left.Left = &treenode2.TreeNode{Val: 4}
	root.Right = &treenode2.TreeNode{Val: 3}
	str := tree2str2(root)
	fmt.Println("str is ", str)
}

/*
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
示例 1:
输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4
输出: "1(2(4))(3)"
解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:
输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4
输出: "1(2()(4))(3)"
解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
*/
func tree2str(t *treenode2.TreeNode) string {
	if t == nil {
		return ""
	}
	s := dfs(t, "")
	return string([]byte(s)[1 : len(s)-1])
}

func dfs(root *treenode2.TreeNode, s string) string {
	if root == nil {
		return ""
	}
	left := dfs(root.Left, s)
	right := dfs(root.Right, s)

	if right != "" {
		if left == "" {
			s = "(" + strconv.Itoa(root.Val) + "()" + right + ")"
		} else {
			s = "(" + strconv.Itoa(root.Val) + left + right + ")"
		}
	} else {
		s = "(" + strconv.Itoa(root.Val) + left + ")"
	}
	return s
}
