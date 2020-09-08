package main

import (
	"fmt"
)

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
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
	dp := make([][]bool, 0)
	for i := 0; i <= len(s1); i++ {
		dp = append(dp, make([]bool, len(s2)))
	}
	if s1[0] == s3[0] {
		dp[1][0] = true
	}
	if s2[0] == s3[0] {
		dp[0][1] = true
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; i <= len(s2); j++ {
			if s2[i+j-1] == s1[i] {
				dp[i][j] = dp[i-1][j]
			} else if s3[i+j-1] == s2[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				return false
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
