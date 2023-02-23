package main

func main() {

}

/*
给你一个二叉树的根结点，请你找出出现次数最多的子树元素和。一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
你需要返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
示例 1：
输入:

	 5
	/  \

2   -3
返回 [2, -3, 4]，所有的值均只出现一次，以任意顺序返回所有值。
示例 2：
输入：

	 5
	/  \

2   -5
返回 [2]，只有 2 出现两次，-5 只出现 1 次。
提示： 假设任意子树元素和均可以用 32 位有符号整数表示。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	list := make([]int, 0)
	dfs(root, &list)
	hash := make(map[int]int)
	max := 0
	for _, n := range list {
		hash[n]++
		if hash[n] > max {
			max = hash[n]
		}
	}
	// 取值
	for k, n := range hash {
		if n == max {
			ans = append(ans, k)
		}
	}
	return ans
}

// 这里可以使用map记录，不然会有太多的重复计算
func dfs(root *TreeNode, list *[]int) int {
	if root == nil {
		return 0
	}
	sum := dfs(root.Left, list) + dfs(root.Right, list) + root.Val
	*list = append(*list, sum)
	return sum
}
