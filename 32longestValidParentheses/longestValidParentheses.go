package main

import (
	"fmt"
)

func main() {
	ans := longestValidParentheses("))((()()()")
	fmt.Println("ans is ", ans)
}

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
*/
func longestValidParentheses(s string) int {
	ss, ans, stark := []byte(s), 0, []int{-1}
	for i, b := range ss {
		if b == '(' {
			stark = append(stark, i)
		} else {
			if len(stark) > 0 {
				stark = stark[:len(stark)-1]
			}
			if len(stark) == 0 {
				stark = append(stark, i)
				continue
			}
			if i-stark[len(stark)-1] > ans {
				ans = i - stark[len(stark)-1]
			}
		}
	}
	return ans
}
