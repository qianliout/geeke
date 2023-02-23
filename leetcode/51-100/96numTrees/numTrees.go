package main

func main() {

}

/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1 // nil节点
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 这里要好好理解，意思是选j这个值做为node结点的值，那么left就有dp[j-1],right就有dp[i-j]个
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}
