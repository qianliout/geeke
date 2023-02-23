package main

import (
	"fmt"
)

func main() {
	res := isMatch("aa", "a")
	fmt.Println(res)
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
示例 1：
输入：s = "aa" p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:
输入：s = "aa" p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：
输入：s = "ab" p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4：
输入：s = "aab" p = "c*a*b"
输出：true
解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5：
输入：s = "mississippi" p = "mis*is*p*."
输出：false
*/
func isMatch(s string, p string) bool {
	ss, pp := []byte(s), []byte(p)
	dp := make(map[int]map[int]bool)
	// dp[i:j]表示 s[:i+1] 和p[:j+1]是否匹配
	// 初值
	for i := 0; i <= len(ss); i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true

	for i := 1; i <= len(ss); i++ {
		for j := 1; j <= len(pp); j++ {
			if dp[i][j-1] && pp[j-1] == '.' {
				dp[i][j] = true
			} else if (dp[i][j-1] && pp[j-1] == '*') || (dp[i-1][j] && pp[j-1] == '*') {
				dp[i][j] = true
			} else if dp[i-1][j-1] && ss[i-1] == pp[j-1] {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[len(s)][len(p)]
}
