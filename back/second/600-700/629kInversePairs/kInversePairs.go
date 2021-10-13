package main

func main() {

}

/*
给出两个整数n和k，找出所有包含从1到n的数字，且恰好拥有k个逆序对的不同的数组的个数。
逆序对的定义如下：对于数组的第i个和第j个元素，如果满i<j且a[i]>a[j]，则其为一个逆序对；否则不是。
由于答案可能很大，只需要返回 答案 mod 109+ 7 的值。
示例 1:
输入: n = 3, k = 0
输出: 1
解释:
只有数组 [1,2,3] 包含了从1到3的整数并且正好拥有 0 个逆序对。
示例 2:
输入: n = 3, k = 1
输出: 2
解释:
数组 [1,3,2] 和 [2,1,3] 都有 1 个逆序对。
*/
func kInversePairs(n int, k int) int {
	/*
		一、状态定义
		int[][] dp = new int[n + 1][k + 1]
		dp[i][j] 表示数字 [1 .. i] 的排列中恰好包含 j 个逆序对的数组的个数

		二、转移方程
		dp[i][j] = dp[i][j - 1] + dp[i - 1][j] - dp[i - 1][j - i]
		转移方程的来源可以参考官方题解

		三、初始化
		dp[i][0] = 1，0个逆序数对只有一种情况

		四、考虑输出
		dp[n][k]
	*/
	m := 1000000007
	dp := make(map[int]map[int]int)
	for i := 0; i <= n; i++ {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k && j <= i*(i-1)/2; j++ {
			// dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-i]
			v := dp[i-1][j] + m
			if j-i >= 0 {
				v = v - dp[i-1][j-i]
			}
			v = v % m
			dp[i][j] = (dp[i][j-1] + v) % m

		}
	}
	return dp[n][k]
}
