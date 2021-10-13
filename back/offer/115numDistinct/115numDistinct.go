package main

import (
	"fmt"
)

func main() {
	hello := make(map[string]interface{})
	hello["3"] = []string{"hello", "owrd"}
	hello["3"] = 3

	fmt.Println(hello)
}

/*
  给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
  字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
  题目数据保证答案符合 32 位带符号整数范围。
  输入：s = "rabbbit", t = "rabbit"
  输出：3
*/

func numDistinct(s string, t string) int {
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	dp := make(map[int]map[int]int)
	for i := 0; i <= len(t); i++ {
		dp[i] = make(map[int]int)
	}
	// 初值,这里的初值很不容易理解
	for j := 0; j <= len(s); j++ {
		dp[0][j] = 1
	}
	// 状态转移
	for i := 0; i <= len(t); i++ {
		for j := 0; j <= len(s); j++ {
			if t[i] == s[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
