package main

import (
	"fmt"
	"math"
)

func main() {
	s := "abaaaab"
	res := longestPalindrome2(s)
	fmt.Println("res is ", res)
}

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"
*/

// 使用动态规划
func longestPalindrome(s string) string {
	var maxLenght int = math.MinInt64
	min, max := 0, 0

	byteSlice := []byte(s)

	// 表示[i:j]是不是回文,注意这里的左右都包含
	dp := make(map[int]map[int]bool)
	// 初值
	for i := 0; i < len(byteSlice); i++ {
		dp[i] = map[int]bool{i: true}
		if i+1 < len(byteSlice) && byteSlice[i+1] == byteSlice[i] {
			dp[i][i+1] = true
		}
	}

	for i := 0; i < len(byteSlice); i++ {
		for j := 0; j < i; j++ {
			if dp[i-1][j+1] && byteSlice[i] == byteSlice[j] {
				dp[i][j] = true
				if i-j > maxLenght {
					min, max = j, i
					maxLenght = i - j
				}
			}
		}
	}
	return string(byteSlice[min : max+1])
}

// 使用双指针
func longestPalindrome2(s string) string {
	if len(s) <= 0 {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter([]byte(s), i, i)
		left2, right2 := expandAroundCenter([]byte(s), i, i+1)

		if right1-left1 > right2-left2 {
			if right1-left1 > end-start {
				start = left1
				end = right1
			}
		} else {
			if right2-left2 > end-start {
				start = left2
				end = right2
			}
		}
	}
	return string([]byte(s)[start : end+1])
}

func expandAroundCenter(s []byte, start, end int) (int, int) {
	left, right := start, end
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 因为不符条件时,left已加1,right已加1,所以这里的返回值要注意
	return left + 1, right - 1
}
