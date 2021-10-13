package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	res := findMode(root)
	fmt.Println("res is ", res)
}

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
假定 BST 有如下定义：
    结点左子树中所含结点的值小于等于当前结点的值
    结点右子树中所含结点的值大于等于当前结点的值
    左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],
   1
    \
     2
    /
   2
返回[2].
提示：如果众数超过1个，不需考虑输出顺序
进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	list := make([]int, 0)
	inOrder(root, &list)

	fmt.Println("list is ", list)

	max, start, end := 1, 0, 1
	res = append(res, list[0])

	for start < len(list) {
		for end < len(list) && list[end] == list[start] {
			end++
		}
		if end-start+1 > max {
			max = end - start + 1
			res = []int{list[start]}
		} else if end-start+1 == max {
			res = append(res, list[start])
		}
		start = end
	}

	return res
}

// 方法一,中序遍历，然后就是一个递增的数组，之后就可以求了
func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
