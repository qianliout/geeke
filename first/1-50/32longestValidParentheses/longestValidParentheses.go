package main

import (
	"fmt"
	"math"
)

func main() {
	s := ")()())()()()"

	res := longestValidParentheses(s)
	fmt.Println("res is ", res)

}

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

// 可以使用satck这种数据结构
func longestValidParentheses(s string) int {
	res := 0
	if len(s) <= 1 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, -1)

	slice := []byte(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] == '(' {
			stack = append(stack, i)
		} else if slice[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = int(math.Max(float64(res), float64(i-stack[len(stack)-1])))
			}
		}
	}
	return res
}
