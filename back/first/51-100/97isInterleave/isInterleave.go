package main

import (
	"fmt"
)

func main() {
	s1 := "aa"
	s2 := "ab"
	s3 := "abaa"
	res := isInterleave(s1, s2, s3)
	fmt.Println("res is ", res)
}

/*
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
示例 1：
输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：
输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if s1 == "" {
		return s2 == s3
	} else if s2 == "" {
		return s1 == s3
	}

	// dp[i][j]表示s1[:i],s2[:j]是否正确,前包后不包
	// 初始化dp
	dp := make([][]bool, 0)
	for i := 0; i <= len(s1)+1; i++ {
		dp = append(dp, make([]bool, len(s2)+1))
	}
	// 初值 这里赋初值很重要，不然不能得到正确答案
	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= len(s2); i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			// 状态转移方程，s3[i+j-1]要么等于s1[i-1],要么等于s2[j-1],
			dp[i][j] = (dp[i-1][j] && s3[i+j-1] == s1[i-1]) || (dp[i][j-1] && s3[i+j-1] == s2[j-1])
		}
	}
	return dp[len(s1)][len(s2)]
}
