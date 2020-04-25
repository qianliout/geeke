package main

import (
	"fmt"
)

func main() {
	res := Ismath("hhhhhh", "h*")
	fmt.Println("res is ", res)
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
说明:
    s 可能为空，且只包含从 a-z 的小写字母。
    p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:


输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false
*/

func isMatch(s string, p string) bool {
	ss := []byte(s)
	pp := []byte(p)
	dp := make(map[int]map[int]bool)
	sn := len(ss)
	pn := len(pp)
	for i := 0; i <= sn; i++ {
		dp[i] = make(map[int]bool)
	}
	// 初值
	dp[0][0] = true

	// 状态转移方程
	for i := 1; i < sn; i++ {
		for j := 1; j < pn; j++ {
			if ss[i-1] == pp[j-1] || pp[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if pp[j] == '.' {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = false
			}

		}
	}
	return dp[sn][pn]
}

func Ismath(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	first := len(s) > 0 && ([]byte(s)[0] == []byte(p)[0] || []byte(p)[0] == '.')
	if len(p) >= 2 && []byte(p)[1] == '*' {
		return Ismath(s, string([]byte(p)[2:])) || (first && Ismath(string([]byte(s)[1:]), p))
	}
	return first && Ismath(string([]byte(s)[1:]), string([]byte(p)[1:]))
}
