package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	w1, w2 := "horse", "ros"
	distance := minDistance(w1, w2)
	fmt.Println("distance is ", distance)
}

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符
示例 1：
*/
func minDistance(word1 string, word2 string) int {
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}

	ww1, ww2 := []byte(word1), []byte(word2)
	dp := make([][]int, len(ww1)+1)
	for i := range dp {
		dp[i] = make([]int, len(ww2)+1)
	}
	// 初值
	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 题目是把word1变成word2,并不是说两个单词相同就好
				dp[i][j] = Min(
					dp[i-1][j],   // 删除word1[i] 使 word1[0:i-1]==word2[0:j],可以删除，因为之后可以再增加
					dp[i][j-1],   // 插入word1[i] 使用 word[0:i]==word2[0:j-1] 可以只等于word2[0:j-1]因为之后可以再增加
					dp[i-1][j-1], // 替换word[i] 使用word[0:i]==word2[0:j-1]
				) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
