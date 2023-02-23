package main

import (
	"fmt"

	"qianliout/leetcode/leetcode/common/utils"
)

func main() {
	fmt.Println(minDistance("seaa", "eat"))
	fmt.Println(minDistance("a", "b"))
}

/*
给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
示例：
输入: "sea", "eat"
输出: 2
解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
提示：
给定单词的长度不超过500。
给定单词中的字符只含有小写字母。
*/

// 最大最小，一定是dp
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	dp := make(map[int]map[int]int)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make(map[int]int)
		dp[i][0] = i // 初值
	}
	// 初值
	for i := 0; i <= len(word2); i++ {
		dp[0][i] = i
	}
	w1, w2 := []byte(word1), []byte(word2)
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if w1[i-1] == w2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.Min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
