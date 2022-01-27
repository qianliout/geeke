package main

import (
	"fmt"

	"qianliout/leetcode/back/common"
)

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
}

/*

给定两个字符串s1, s2，找到使两个字符串相等所需删除字符的ASCII值的最小和。
 示例 1:
输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
 示例 2:
输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。
 注意:
 0 < s1.length, s2.length <= 1000。
 所有字符串中的字符ASCII值在[97, 122]之间。
*/

func minimumDeleteSum(s1 string, s2 string) int {
	ss1, ss2 := []byte(s1), []byte(s2)
	// dp[i][j] 表示ss[:i][:j]的值，其中不包含i和j
	dp := make(map[int]map[int]int)
	// 初始化
	sum1, sum2 := 0, 0
	for i := 0; i <= len(ss1); i++ {
		dp[i] = make(map[int]int)
	}
	for i := 1; i <= len(ss1); i++ {
		sum1 += int(ss1[i-1])
		dp[i][0] = sum1
	}

	for i := 1; i <= len(ss2); i++ {
		sum2 += int(ss2[i-1])
		dp[0][i] = sum2
	}

	// dp[0][0] = math.MaxInt64
	for i := 1; i <= len(ss1); i++ {
		for j := 1; j <= len(ss2); j++ {
			if ss1[i-1] == ss2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = common.Min(dp[i-1][j]+int(ss1[i-1]),
					dp[i][j-1]+int(ss2[j-1]),
					dp[i-1][j-1]+int(ss1[i-1])+int(ss2[j-1]))
			}
		}
	}
	return dp[len(ss1)][len(ss2)]
}
