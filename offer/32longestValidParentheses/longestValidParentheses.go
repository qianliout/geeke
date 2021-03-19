package main

import (
	"fmt"
)

func main() {
	ans := longestValidParentheses("()())")
	fmt.Println(ans)
}

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
*/
func longestValidParentheses(s string) int {
	ss := []byte(s)
	stark := make([]int, 0)
	stark = append(stark, -1)
	ans := 0
	for i, v := range ss {
		if v == '(' {
			stark = append(stark, i)
		} else {
			stark = stark[:len(stark)-1]
			if len(stark) == 0 {
				stark = append(stark, i)
			} else {
				if ans < i-stark[len(stark)-1] {
					ans = i - stark[len(stark)-1]
				}
			}
		}
	}
	return ans
}
