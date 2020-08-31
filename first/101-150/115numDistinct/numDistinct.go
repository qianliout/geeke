package main

import (
	"fmt"
)

func main() {
	// res := numDistinct("rabbbit", "rabbit")
	res := numDistinct("babgbag", "bag")
	fmt.Println("res is ", res)
}

/*
给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
题目数据保证答案符合 32 位带符号整数范围。
示例 1：
输入：S = "rabbbit", T = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
示例 2：
输入：S = "babgbag", T = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。
(上箭头符号 ^ 表示选取的字母)
babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
*/

func numDistinct(s string, t string) int {
	/*
		动态规划
		dp[i][j] 代表 T 前 i 字符串可以由 S j 字符串组成最多个数.
		所以动态方程:
		当 S[j] == T[i] , dp[i][j] = dp[i-1][j-1] + dp[i][j-1];
		当 S[j] != T[i] , dp[i][j] = dp[i][j-1]
		我相信有一部分人不懂是 当s[j] == t[i] 时的情况的分析
		假设我们有字符串 s = abcc 和 字符串 abc
		设定dp[i][j]为使用s的前i个字符能够最多组成多少个t的前j个字符
		当s[j] == t[i] 时：此时假设j=3，i=2，需要比较的是 abcc 中含有多少个abc
		当考虑使用第i个元素时，即我们让 abcc 中的最后一个c 和abc最后一个c匹配上，这时我们需要看的是 abc中含有多少个 ab ,对应dp[i-1][j-1]
		2.当不考虑使用第i个元素时，我们需要看的是 abc 中含有多少个abc ,对应dp[i][j-1]
		转移方程为dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
		当s[j] ！= t[i] 时，此时假设j=3，i=2
		只有一种情况，只需要看的是 abc 中含有多少个abc即可，对应dp[i][j-1]
		转移方程为dp[i][j] = dp[i][j-1]
	*/
	if len(s) == 0 || len(t) == 0 {
		return 0
	}
	// dp[i][j] 代表 T 前 i 字符串可以由 S j 字符串组成最多个数. 注意这里是不包括 i 和 j的
	dp := make(map[int]map[int]int)
	// 初始化
	for i := 0; i <= len(t); i++ {
		dp[i] = make(map[int]int)
	}
	// 初值
	for j := 0; j <= len(s); j++ {
		dp[0][j] = 1
	}
	
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
