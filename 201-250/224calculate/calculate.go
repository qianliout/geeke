package main

import (
	"fmt"
	"strconv"
)

func main() {
	calculate("(1+(4+5+20)-3)+(6+48)")
}

/*
实现一个基本的计算器来计算一个简单的字符串表达式的值。
字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。
示例 1:
输入: "1 + 1"
输出: 2
示例 2:
输入: " 2-1 + 2 "
输出: 3
示例 3:
输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
*/
func calculate(s string) int {
	symbolStark := make([]string, 0)
	numStark := make([]int, 0)
	if len(s) == 0 {
		return 0
	}
	ss := []byte(s)
	length := len(ss)
	symbolMap := map[string]bool{"(": true, ")": true, "+": true, "-": true}
	i := 0

	for i < length {
		if symbolMap[string(ss[i])] {
			symbolStark = append(symbolStark, string(ss[i]))
		} else {
			num, _ := strconv.Atoi(string(ss[i]))
			for !symbolMap[string(ss[i+1])] {
				r, _ := strconv.Atoi(string(ss[i+1]))
				num = num*10 + r
				i++
			}
			numStark = append(numStark, num)
		}
		i++
	}
	first := false

	symbolStark2 := make([]string, 0)
	numStark2 := make([]int, 0)
	if symbolStark[0] == "(" {
		first = true
	}
	if first {
		for len(symbolStark) > 0 && len(numStark) > 0 {
			symbol := symbolStark[0]
			symbolStark = symbolStark[1:]
			if symbol == "(" {
				symbolStark2 = append(symbolStark2, symbol)
			}
			if symbol

		}

	}

	fmt.Println(symbolStark, numStark)
	return 0
}
