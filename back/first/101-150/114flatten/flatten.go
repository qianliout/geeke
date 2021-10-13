package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {
	root := treenode.TreeNode{Val: 0}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Left = &TreeNode{Val: 3}
	// root.Left.Right = &TreeNode{Val: 4}
	//
	// root.Right = &TreeNode{Val: 5}
	// root.Right.Right = &TreeNode{Val: 8}
	flatten3(&root)
	treenode.PreOrderTraversal(&root)
}

/*
给定一个二叉树，原地将它展开为链表。
例如，给定二叉树
    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
*/
// 使用递归的方法是最容易想到的，只是几个条件判断要注意
func flatten(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		flatten(root.Left)
	}
	if root.Right != nil {
		flatten(root.Right)
	}
	var pre *treenode.TreeNode
	if root.Left != nil {
		pre = root.Left
		lastLeft := pre.Right
		for lastLeft != nil {
			pre = lastLeft
			lastLeft = lastLeft.Right
		}
	}

	preRight := root.Right
	if root.Left != nil {
		root.Right = root.Left
	}
	root.Left = nil
	if pre != nil {
		pre.Right = preRight
	}
}

// 使用迭代的就去，这个方法不容易想到
/*
   将左子树插入到右子树的地方
   将原来的右子树接到左子树的最右边节点
   考虑新的右子树的根节点，一直重复上边的过程，直到新的右子树为 null
	这种解法应该形成自己的模板
*/
func flatten2(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
}

var pre *treenode.TreeNode

func flatten3(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	flatten3(root.Right)
	flatten3(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}

func flatten4(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	for root != nil {
		if root.Left != nil {
			// 找到左子树的最小面的那个右子树
			lastRight := root.Left
			for lastRight.Right != nil {
				lastRight = lastRight.Right
			}
			// 把原来的right接到现在的这个最小面的右子树上
			lastRight.Right = root.Right
			// 把原来的left放到root的右边
			root.Right = root.Left
			// 左边置空
			root.Left = nil
		}
		// 判断下一个root
		root = root.Right
	}
}
