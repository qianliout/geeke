package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := calculate("-3+5 / 2")
	fmt.Println("res is ", res)
}

/*
实现一个基本的计算器来计算一个简单的字符串表达式的值。
字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
示例 1:
输入: "3+2*2"
输: 7
示例 2:
输入: " 3/2 "
输出: 1
示例 3:
输入: " 3+5 / 2 "
输出: 5
说明：
    你可以假设所给定的表达式都是有效的。
    请不要使用内置的库函数 eval。
*/

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	numStark := make([]int, 0)
	ss := []byte(s)
	num := 0
	var sign byte
	sign = '+'
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(string(s[i]))
		if err == nil || i == len(s) {
			num = num*10 + n
		}
		if i == len(s)-1 || ss[i] == '+' || ss[i] == '-' || ss[i] == '*' || ss[i] == '/' {
			
			if sign == '+' {
				numStark = append(numStark, num)
			} else if sign == '-' {
				numStark = append(numStark, -num)
			} else if sign == '*' {
				last := numStark[len(numStark)-1]
				numStark = numStark[:len(numStark)-1]
				numStark = append(numStark, last*num)
			} else if sign == '/' {
				last := numStark[len(numStark)-1]
				numStark = numStark[:len(numStark)-1]
				numStark = append(numStark, last/num)
			}
			num, sign = 0, ss[i]
		}
	}
	sum := 0
	for _, n := range numStark {
		sum += n
	}
	return sum
}
