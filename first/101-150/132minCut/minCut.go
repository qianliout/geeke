package main

import (
	"fmt"
)

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的最少分割次数。
示例:
输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
*/

func main() {
	s := "a"
	res := minCut(s)
	fmt.Println("res is ", res)
}

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}
	ss := []byte(s)
	// 	 dp[i]表示s[0:i+1]这个下串最段分隔次数,注意这里是包含了i这个
	dp := make(map[int]int)
	//  初值,都分成单个词，那肯定都是回文串了
	for i := 0; i < len(s); i++ {
		dp[i] = i
	}
	// 状态转移方程
	for i := 1; i < len(ss); i++ {
		if verifyThatItIsAPalindromeString(ss, 0, i) {
			dp[i] = 0
			continue
		}
		
		for j := i-1; j >= 0; j-- {
			if verifyThatItIsAPalindromeString(ss, j+1, i) {
				if dp[i] > dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[len(s)-1]
}

func verifyThatItIsAPalindromeString(s []byte, start, end int) bool {
	for end > start {
		if s[end] != s[start] {
			return false
		}
		end--
		start++
	}
	return true
}
