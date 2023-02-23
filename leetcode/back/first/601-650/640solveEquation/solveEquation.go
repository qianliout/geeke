package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solveEquation("x+5-3+x=6+x-2"))
	fmt.Println(solveEquation("x+2x=3x"))
	fmt.Println(solveEquation("2x+3x-6x=x+2"))
	fmt.Println(solveEquation("-2x=2"))
}

/*
求解一个给定的方程，将x以字符串"x=#value"的形式返回。该方程仅包含'+'，' - '操作，变量 x 和其对应系数。
如果方程没有解，请返回“No solution”。
如果方程有无限解，则返回“Infinite solutions”。
如果方程中只有一个解，要保证返回值 x 是一个整数。
示例 1：
输入: "x+5-3+x=6+x-2"
输出: "x=2"
示例 2:
输入: "x=x"
输出: "Infinite solutions"
示例 3:
输入: "2x=x"
输出: "x=0"
示例 4:
输入: "2x+3x-6x=x+2"
输出: "x=-1"
示例 5:
输入: "x=x+2"
输出: "No solution"
*/

// 这道题就是考察字符串的用法
func solveEquation(equation string) string {
	sx := 0
	num := 0
	pre := 0
	ss := []byte(equation)
	op := 1
	reverse := 1
	for i, v := range ss {
		if v >= '0' && v <= '9' {
			pre = pre*10 + int(v-'0')
			// 最后可能是一个数字 //这一步容易出错
			if i == len(ss)-1 {
				num += pre * op * reverse
			}
		} else if v == '-' {
			num += pre * op * reverse
			op = -1
			pre = 0
		} else if v == '+' {
			num += pre * op * reverse
			op = 1
			pre = 0
		} else if v == '=' {
			num += pre * op * reverse
			reverse = -1
			op = 1 // 这一步容易出错
			pre = 0
		} else if v == 'x' {
			if i-1 < 0 || ss[i-1] == '+' || ss[i-1] == '-' || ss[i-1] == '=' {
				pre = 1
			}
			sx += pre * op * reverse
			pre = 0
		}
	}

	// 根据结果返回值了
	if sx == 0 && num == 0 {
		return "Infinite solutions"
	}
	if sx == 0 && num != 0 {
		return "No solution"
	}

	return "x=" + strconv.Itoa(-num/sx)
}
