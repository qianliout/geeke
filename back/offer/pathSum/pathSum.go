package main

import (
	treenode2 "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
给定一个二叉树，它的每个结点都存放一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1
返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
*/

// func pathSum(root *TreeNode, sum int) int {
//
// 	visit := make(map[*TreeNode]bool)
// 	var res int
// 	helper(root, sum, 0, &res, visit)
// 	return res
// }

func pathSum2(root *treenode2.TreeNode, sum int) int {
	var res int
	visit := make(map[*treenode2.TreeNode]bool)
	helper(root, sum, 0, &res, visit)

	return res
}

func helper(root *treenode2.TreeNode, sum, path int, res *int, visit map[*treenode2.TreeNode]bool) {
	if root == nil {
		return
	}
	if path == sum {
		*res += 1
	}
	visit[root] = true
	// 走左边
	helper(root.Left, sum, path-root.Val, res, visit)
	helper(root.Left, sum, path, res, visit)
	// 走右边
	helper(root.Left, sum, path-root.Val, res, visit)
	helper(root.Left, sum, path, res, visit)
	visit[root] = false
}

func pathSum(root *treenode2.TreeNode, sum int) int {
	// sumList := make([]int, 0)
	// n := dfs(root, &sumList, sum)
	// return n

	prefixSumCount := make(map[int]int)
	// 初值,前缀和为0时有一种方法，就是什么结点都不选
	prefixSumCount[0] = 1
	return recursionPathSum(root, prefixSumCount, sum, 0)
}

// 前缀和的思想
func recursionPathSum(node *treenode2.TreeNode, prefixSumCount map[int]int, target, currSum int) int {
	// 1.递归终止条件
	if node == nil {
		return 0
	}
	// 本层要做的事
	res := 1
	// 当前路径上的和
	currSum += node.Val
	// 看看root到当前节点这条路上是否存在节点前缀和加target为currSum的路径
	// 当前节点->root节点反推，有且仅有一条路径(只能父到子结点，不能 左->root->右)，如果此前有和为currSum-target,而当前的和又为currSum,两者的差就肯定为target了
	// currSum-target相当于找路径的起点，起点的sum+target=currSum，当前点到起点的距离就是target
	res += prefixSumCount[currSum-target]
	prefixSumCount[currSum] += 1
	// 进入下一层
	res += recursionPathSum(node.Left, prefixSumCount, target, currSum)
	res += recursionPathSum(node.Right, prefixSumCount, target, currSum)
	// 回到本层把进入下一层的影响消掉
	prefixSumCount[currSum] -= 1
	return res
}

func dfs(root *treenode2.TreeNode, sumList *[]int, sum int) int {
	if root == nil {
		return 0
	}
	// 这里一定要进行赋值操作，不然就不能得到正确的结果
	nl := append([]int{}, *sumList...)

	for i, n := range nl {
		nl[i] = n + root.Val
	}
	nl = append(nl, root.Val)
	count := 0
	for _, n := range nl {
		if n == sum {
			count++
		}
	}
	count += dfs(root.Left, &nl, sum)
	count += dfs(root.Right, &nl, sum)

	return count
}
