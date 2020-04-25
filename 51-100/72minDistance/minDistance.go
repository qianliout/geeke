package main

import (
	"fmt"

	. "outback/leetcode/common"
)

func main() {
	res := minDistance("", "")
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
	// 这个判断不写也可以,写了也不报错
	//if len(word1) == 0 {
	//	return len(word2)
	//}
	//if len(word2) == 0 {
	//	return len(word1)
	//}

	dp := make(map[int]map[int]int)
	ww1 := []byte(word1)
	ww2 := []byte(word2)
	n1 := len(word1)
	n2 := len(word2)

	for i := 0; i <= n1; i++ {
		dp[i] = make(map[int]int)
	}
	// 初值
	for j := 1; j <= n2; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}

	for j := 1; j <= n1; j++ {
		dp[j][0] = dp[j-1][0] + 1
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if ww1[i-1] == ww2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}

		}
	}
	return dp[n1][n2]
}
