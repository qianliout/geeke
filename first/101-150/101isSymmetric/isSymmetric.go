package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 3}
	res := bfs(&root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
链接：https://leetcode-cn.com/problems/symmetric-tree
*/
// bfs
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		thisLeve := make([]*TreeNode, 0)
		for _, n := range queue {

			if n == nil {
				continue
			}
			if n.Left != nil {
				thisLeve = append(thisLeve, n.Left)
			} else {
				thisLeve = append(thisLeve, nil)
			}
			if n.Right != nil {
				thisLeve = append(thisLeve, n.Right)
			} else {
				thisLeve = append(thisLeve, nil)
			}
		}
		if !check(thisLeve) {
			return false
		}
		queue = thisLeve
	}
	return true
}

func check(queue []*TreeNode) bool {
	if len(queue)&1 == 1 {
		return false
	}
	for i := 0; i < len(queue)/2; i++ {
		f := queue[i]
		s := queue[len(queue)-i-1]
		if f == nil && s == nil {
			continue
		} else if f == nil || s == nil {
			return false
		} else if f.Val != s.Val {
			return false
		}
	}
	return true
}

// dfs
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetric(left.Right, right.Left) && symmetric(left.Left, right.Right)
}

func dfs(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	// 这一步是一个技巧，因为前面判断了都是nil,这里就不可能都是nil
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func bfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		thisLevel := make([]*TreeNode, 0)
		for _, node := range queue {
			if node == nil {
				continue
			}
			thisLevel = append(thisLevel, node.Left, node.Right)
		}
		if !check2(thisLevel) {
			return false
		}
		queue = thisLevel
	}
	return true
}

func check2(queue []*TreeNode) bool {
	left, right := 0, len(queue)-1
	for left < right {
		if queue[left] == nil && queue[right] == nil {
			left++
			right--
			continue
		}
		if queue[left] == nil || queue[right] == nil {
			return false
		}
		if queue[left].Val != queue[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}
