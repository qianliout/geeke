package main

import (
	"fmt"
)

func main() {
	s := "b"
	res := longestPalindrome(s)
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

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)

		if r1-l1 > end-start  {
			start = l1
			end = r1
		}
		if r2-l2 > end-start {
			start = l2
			end = r2
		}
	}

	return string([]byte(s)[start : end+1])
}

// 以start,end为中心向两边扩散
func expandAroundCenter(s string, start, end int) (int, int) {
	left, right := start, end
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
