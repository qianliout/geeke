package main

import (
	"fmt"

	"qianliout/leetcode/common/utils"
)

func main() {
	word1 := "horse"
	word2 := "ros"
	res := minDistance(word1, word2)
	fmt.Println("res is ", res)
}

/*
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
你可以对一个单词进行如下三种操作：
    插入一个字符
    删除一个字符
    替换一个字符
示例 1：
输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：
输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
*/

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	// 定义dp,dp表示第i,j之前时相同的最小编辑，也就是下标为i-1,i-1
	dp := make(map[int]map[int]int)
	len1 := len(word1)
	len2 := len(word2)

	// 初始化
	for i := 0; i <= len1; i++ {
		dp[i] = make(map[int]int)
	}
	// 初值
	for i := 1; i <= len1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}
	for j := 1; j <= len2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	// 状态转移方程
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len1][len2]
}

func minDistance2(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	dp := make(map[int]map[int]int)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make(map[int]int)
		dp[i][0] = i
	}
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}

	// 状态转移
	for i := 1; i <= len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i-1][j-1] 表示替换操作，dp[i-1][j] 表示删除操作，dp[i][j-1] 表示插入操作
				dp[i][j] = utils.Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
