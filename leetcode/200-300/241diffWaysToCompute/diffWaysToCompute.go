package main

import (
	"fmt"
	"strconv"
)

func main() {
	expression := "-2"
	compute := diffWaysToCompute(expression)
	fmt.Println(compute)
}

/*
给你一个由数字和运算符组成的字符串 expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。
生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 104 。
输入：expression = "2*3-4*5"
输出：[-34,-14,-10,-10,10]
解释：
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
*/
func diffWaysToCompute(expression string) []int {
	if n, err := strconv.Atoi(expression); err == nil {
		return []int{n}
	}
	ss := []byte(expression)
	ans := make([]int, 0)
	for i := range ss {
		if ss[i] >= '0' && ss[i] <= '9' {
			continue
		}
		left := diffWaysToCompute(string(ss[:i]))
		right := diffWaysToCompute(string(ss[i+1:]))
		for j := range left {
			for k := range right {
				switch ss[i] {
				case '+':
					ans = append(ans, left[j]+right[k])
				case '-':
					ans = append(ans, left[j]-right[k])
				case '*':
					ans = append(ans, left[j]*right[k])
				case '/':
					ans = append(ans, left[j]/right[k])
				}
			}
		}
	}
	return ans
}

func Digit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
