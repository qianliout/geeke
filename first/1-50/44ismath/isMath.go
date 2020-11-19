package main

import (
	"fmt"
)

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。
说明:
    s 可能为空，且只包含从 a-z 的小写字母。
    p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:
输入:
s = "acdcb"
p = "a*c?b"
输入: false
*/

func main() {
	// res := isMathisMatch("acdcb", "a*c?b")
	// res := isMathisMatch("aa", "*")
	res := isMathisMatch("adceb", "*a*b")
	fmt.Println(res)
}

func isMathisMatch2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	ss := []byte(s)
	pp := []byte(p)

	i, j, iStart, jStart := 0, 0, -1, -1
	for i < len(ss) {
		if j < len(pp) && (ss[i] == pp[j] || pp[j] == '?') {
			i++
			j++
		} else if j < len(pp) && pp[j] == '*' {
			iStart = i
			jStart = j
			j++
		} else if iStart > -1 {
			iStart = iStart + 1
			i = iStart
			j = jStart + 1
		} else {
			return false
		}
	}
	for j < len(pp) && pp[j] == '*' {
		j++
	}
	return j == len(pp)
}

// dp才是正解，上面的贪心算法虽然也能得出答案，但是有点不好理解。
func isMathisMatch(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	// dp[i][j]表示，s前i个字符对应p的j个字符的结果
	dp := make(map[int]map[int]bool)
	for i := 0; i <= len(s); i++ {
		dp[i] = make(map[int]bool)
	}
	dp[0][0] = true
	ss := []byte(s)
	pp := []byte(p)

	for j := 1; j <= len(p); j++ {
		dp[0][j] = pp[j-1] == '*' && dp[0][j-1]
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if ss[i-1] == pp[j-1] || pp[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if pp[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
