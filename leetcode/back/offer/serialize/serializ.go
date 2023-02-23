package main

import (
	"strconv"
	"strings"

	treenode2 "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

// dfs序列化
func dfss(root *treenode2.TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + dfss(root.Left) + "," + dfss(root.Right)
}

func dfsd(s string) *treenode2.TreeNode {
	split := strings.Split(s, ",")
	return dfsdh(&split)
}

func dfsdh(ss *[]string) *treenode2.TreeNode {
	if len(*ss) == 0 {
		return nil
	}
	i, err := strconv.Atoi((*ss)[0])
	*ss = (*ss)[1:]
	if err != nil {
		return nil
	}
	root := &treenode2.TreeNode{
		Val:   i,
		Left:  dfsdh(ss),
		Right: dfsdh(ss),
	}
	return root
}

// 使用bfs的方式,就是常用的层序方式
func bfss(root *treenode2.TreeNode) string {
	stark := make([]*treenode2.TreeNode, 0)
	stark = append(stark, root)

	ans := ""
	for len(stark) > 0 {
		first := stark[0]
		stark = stark[1:]
		if first == nil {
			ans = ans + "null" + ","
		} else {
			ans = ans + strconv.Itoa(first.Val) + ","
			stark = append(stark, first.Left)
			stark = append(stark, first.Right)
		}
	}
	return ans
}

func bfsds(s string) *treenode2.TreeNode {
	ss := strings.Split(s, ",")
	if len(ss) == 0 {
		return nil
	}
	idx := 0
	ati, err := strconv.Atoi(ss[idx])
	if err != nil {
		return nil
	}

	root := treenode2.TreeNode{Val: ati}
	idx++
	queue := []*treenode2.TreeNode{&root}

	for len(queue) > 0 {
		nod := queue[0]
		queue = queue[1:]

		if ss[idx] != "null" {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				left := &treenode2.TreeNode{Val: j}
				nod.Left = left
				queue = append(queue, left)
			}
		}
		idx++
		if ss[idx] != "null" {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				right := &treenode2.TreeNode{Val: j}
				nod.Right = right
				queue = append(queue, right)
			}
		}
		idx++
	}
	return &root
}
