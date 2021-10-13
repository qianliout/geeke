package main

import (
	"strconv"
	"strings"

	"outback/leetcode/back/common/treenode"
)

func main() {
	
}

// dfs序列化
func dfss(root *treenode.TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + dfss(root.Left) + "," + dfss(root.Right)
}

func dfsd(s string) *treenode.TreeNode {
	split := strings.Split(s, ",")
	return dfsdh(&split)
}

func dfsdh(ss *[]string) *treenode.TreeNode {
	if len(*ss) == 0 {
		return nil
	}
	i, err := strconv.Atoi((*ss)[0])
	*ss = (*ss)[1:]
	if err != nil {
		return nil
	}
	root := &treenode.TreeNode{
		Val:   i,
		Left:  dfsdh(ss),
		Right: dfsdh(ss),
	}
	return root
}

// 使用bfs的方式,就是常用的层序方式
func bfss(root *treenode.TreeNode) string {
	stark := make([]*treenode.TreeNode, 0)
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

func bfsds(s string) *treenode.TreeNode {
	ss := strings.Split(s, ",")
	if len(ss) == 0 {
		return nil
	}
	idx := 0
	ati, err := strconv.Atoi(ss[idx])
	if err != nil {
		return nil
	}
	
	root := treenode.TreeNode{Val: ati}
	idx++
	queue := []*treenode.TreeNode{&root}
	
	for len(queue) > 0 {
		nod := queue[0]
		queue = queue[1:]
		
		if ss[idx] != "null" {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				left := &treenode.TreeNode{Val: j}
				nod.Left = left
				queue = append(queue, left)
			}
		}
		idx++
		if ss[idx] != "null" {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				right := &treenode.TreeNode{Val: j}
				nod.Right = right
				queue = append(queue, right)
			}
		}
		idx++
	}
	return &root
}
