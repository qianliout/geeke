package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
示例 1:
输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	root := &TreeNode{}
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 != nil && t2 == nil {
		root.Val = t1.Val
		root.Left = mergeTrees(t1.Left, nil)
		root.Right = mergeTrees(t1.Right, nil)
	}
	if t1 == nil && t2 != nil {
		root.Val = t2.Val
		root.Left = mergeTrees(nil, t2.Left)
		root.Right = mergeTrees(nil, t2.Right)
	}
	if t1 != nil && t2 != nil {
		root.Val = t1.Val + t2.Val
		root.Left = mergeTrees(t1.Left, t2.Left)
		root.Right = mergeTrees(t1.Right, t2.Right)
	}
	return root
}
