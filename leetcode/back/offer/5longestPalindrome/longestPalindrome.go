package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome2("cbbd"))
}

/*
给你一个字符串 s，找到 s 中最长的回文子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	ss := []byte(s)
	dp := make(map[int]map[int]bool)
	// 初值
	for i := 0; i < len(ss); i++ {
		dp[i] = make(map[int]bool)
		dp[i][i] = true
		if i+1 < len(ss) && ss[i] == ss[i+1] {
			dp[i][i+1] = true
		}

	}
	ans := string(ss[0])
	for i := 1; i < len(ss); i++ {
		for j := 0; j < i; j++ {
			if dp[i-1][j+1] && ss[i] == ss[j] {
				dp[i][j] = true
				if len(ss[j:i+1]) > len(ans) {
					ans = string(ss[j : i+1])
				}
			}
		}
	}
	return ans
}

// 使用双指针
func longestPalindrome2(s string) string {
	ans := ""

	ss := []byte(s)
	for i := 0; i < len(s); i++ {
		s1, e1 := expansion(ss, i, i)
		s2, e2 := expansion(ss, i, i+1)
		if s1 >= 0 && e1 < len(ss) && len(ss[s1:e1+1]) > len(ans) {
			ans = string(ss[s1 : e1+1])
		}
		if s2 >= 0 && e2 < len(ss) && len(ss[s2:e2+1]) > len(ans) {
			ans = string(ss[s2 : e2+1])
		}
	}
	return ans
}

func expansion(ss []byte, start, end int) (int, int) {
	for start >= 0 && end < len(ss) && ss[start] == ss[end] {
		start--
		end++
	}
	return start + 1, end - 1
}
