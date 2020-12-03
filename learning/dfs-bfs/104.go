package dfsbfs

import (
	"math"
	"outback/leetcode/common/treenode"
)

/*
111. 二叉树的最小深度
104, 二叉树的最大深度
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。
示例:
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

func MaxDepth(root *treenode.TreeNode) int {
	return DfsMaxDepth(root)
}

func DfsMaxDepth(root *treenode.TreeNode) int {
	// dfs解法,一定要理解这种做法,recursion
	if root == nil {
		return 0
	}
	leftDepth := DfsMaxDepth(root.Left)
	rightDepth := DfsMaxDepth(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}

func DfsMaxDepthByStark(root *treenode.TreeNode) int {
	// dfs not usr recursion,but use stark
	if root == nil {
		return 0
	}
	type Result struct {
		Node  *treenode.TreeNode
		Depth int
	}
	depth := 0
	stark := make([]Result, 0)
	stark = append(stark, Result{
		Node:  root,
		Depth: 1,
	})
	for len(stark) > 0 {
		r := stark[len(stark)-1]
		currentDepth := r.Depth
		if r.Node != nil {
			depth = int(math.Max(float64(depth), float64(currentDepth)))
			stark = append(stark, Result{
				Node:  r.Node.Left,
				Depth: currentDepth + 1,
			})
			stark = append(stark, Result{
				Node:  r.Node.Right,
				Depth: currentDepth + 1,
			})
		}
	}
	return depth
}

func BfsMaxDepth(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	max := 0
	for len(queue) > 0 {
		length := len(queue)
		max += 1
		for _, node := range queue {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
	}
	return max
}

func BfsMinDepth(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	min := 0
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		length := len(queue)
		min += 1
		for _, node := range queue {
			if node.Right == nil && node.Left == nil {
				return min
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[length:]
	}
	return min
}

// 最小深度，是根结点到节子结点，所以，一定不能像最大深度那样进行判断
func DfsMinDepth(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depth := math.MaxInt32
	if root.Left != nil {
		depth = int(math.Min(float64(DfsMinDepth(root.Left)), float64(depth)))
	}

	if root.Right != nil {
		depth = int(math.Min(float64(DfsMinDepth(root.Right)), float64(depth)))
	}
	return depth
}

func DfsMinDepthByStark(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	type Result struct {
		Node  *treenode.TreeNode
		Depth int
	}
	stark := make([]Result, 0)
	stark = append(stark, Result{
		Node:  root,
		Depth: 1,
	})
	depth := math.MaxInt32
	for len(stark) > 0 {
		r := stark[len(stark)-1]
		if r.Node != nil { // 因为这里统一判断了，所以在下面加入时可以不判断
			depth = int(math.Min(float64(r.Depth), float64(depth)))
			stark = append(stark, Result{
				Node:  r.Node.Left,
				Depth: depth,
			})
			stark = append(stark, Result{
				Node:  r.Node.Right,
				Depth: depth,
			})
		}
	}
	return depth
}
