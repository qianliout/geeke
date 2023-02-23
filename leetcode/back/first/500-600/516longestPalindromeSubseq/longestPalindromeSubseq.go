package main

import (
	"fmt"

	"qianliout/leetcode/leetcode/common/utils"
)

func main() {
	s := "bbbab"
	res := longestPalindromeSubseq(s)
	fmt.Println("res is ", res)
}

/*
给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
示例 1:
输入:
"bbbab"
输出:
4
一个可能的最长回文子序列为 "bbbb"。
示例 2:
输入:
"cbbd"
输出:
2
一个可能的最长回文子序列为 "bb"。
提示：
1 <= s.length <= 1000
s 只包含小写英文字母
*/

/*
状态
f[i][j] 表示 s 的第 i 个字符到第 j 个字符组成的子串中，最长的回文序列长度是多少。
转移方程
如果 s 的第 i 个字符和第 j 个字符相同的话
f[i][j] = f[i + 1][j - 1] + 2
如果 s 的第 i 个字符和第 j 个字符不同的话
f[i][j] = max(f[i + 1][j], f[i][j - 1])
然后注意遍历顺序，i 从最后一个字符开始往前遍历，j 从 i + 1 开始往后遍历，这样可以保证每个子问题都已经算好了。
初始化
f[i][i] = 1 单个字符的最长回文序列是 1
结果
f[0][n - 1]
*/
func longestPalindromeSubseq(s string) int {
	dp := make(map[int]map[int]int)
	// 初始化加初值
	for i := 0; i <= len(s); i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]int)
		}
		dp[i][i] = 1
	}
	// 计算顺序很重要
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = utils.Max(dp[i+1][j], dp[i][j-1])
			}

		}
	}
	return dp[0][len(s)-1]
}

// func Max(nums ...int) int {
// 	max := math.MinInt64
// 	for _, num := range nums {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	return max
// }
