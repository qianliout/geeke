package main

import (
	"strconv"
)

func main() {

}

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

1 <= s.length <= 3 * 105
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式
'+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
'-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
输入中不存在两个连续的操作符
每个数字和运行的计算将适合于一个有符号的 32位 整数
*/
func calculate(s string) int {
	ss := []byte(s)
	res, num, op := 0, 0, 1
	numStack := make([]int, 0)
	opStack := make([]int, 0)
	for _, ch := range ss {
		if ch == ' ' {
			continue
		}
		if ch == '(' {
			res = res + num*op
			numStack = append(numStack, res)
			opStack = append(opStack, op)
			res, num, op = 0, 0, 1
		} else if ch == ')' {
			res += num * op
			res *= opStack[len(numStack)-1]
			res += numStack[len(numStack)-1]
			num, op = 0, 1
			opStack, numStack = opStack[:len(opStack)-1], numStack[:len(numStack)-1]
		} else if ch == '+' {
			res += num * op
			num, op = 0, 1
		} else if ch == '-' {
			res += num * op
			num, op = 0, -1
		} else {
			n, _ := strconv.Atoi(string(ch))
			num = num*10 + n
		}
	}
	return res + op*num
}
