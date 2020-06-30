package main

import (
	"strconv"
	"strings"
)

func main() {

}

/*
给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。
你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。
示例 1:
输入: "2-1-1"
输出: [0, 2]
解释:
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2:
输入: "2*3-4*5"
输出: [-34, -14, -10, -10, 10]
解释:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/

func diffWaysToCompute(input string) []int {
	if isDigit(input) {
		a, _ := strconv.Atoi(input)
		return []int{a}
	}
	ans := make([]int, 0)
	for i, v := range []byte(input) {
		symbol := string(v)
		if symbol == "+" || symbol == "-" || symbol == "*" {
			left := diffWaysToCompute(string(input[:i]))
			right := diffWaysToCompute(string(input[i+1:]))

			for _, l := range left {
				for _, r := range right {
					tem := 0
					if symbol == "+" {
						tem = l + r
					} else if symbol == "-" {
						tem = l - r
					} else if symbol == "*" {
						tem = l * r
					}
					ans = append(ans, tem)
				}
			}
		}
	}
	return ans
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}

func diffWaysToCompute2(input string) []int {
	res := make([]int, 0)
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 0 {
		return res
	}
	return helper([]byte(input))
}

func helper(input []byte) []int {
	res := make([]int, 0)
	if isDigit(string(input)) {
		n, _ := strconv.Atoi(string(input))
		res = append(res, n)
		return res
	}
	for i, ch := range input {
		if ch == '+' || ch == '*' || ch == '-' {
			left := helper(input[:i])
			right := helper(input[i+1:])
			tem := 0
			for _, l := range left {
				for _, r := range right {
					if ch == '+' {
						tem = l + r
					} else if ch == '-' {
						tem = l - r
					} else if ch == '*' {
						tem = l * r
					}
					res = append(res, tem)
				}
			}
		}
	}
	return res
}


