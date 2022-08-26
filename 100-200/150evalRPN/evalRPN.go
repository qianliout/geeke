package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	val := evalRPN(tokens)
	fmt.Println("var is ", val)
}

/*
根据 逆波兰表示法，求表达式的值。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
注意 两个整数之间的除法只保留整数部分
*/
func evalRPN(tokens []string) int {
	stark := make([]int, 0)
	for i := range tokens {
		if !symbol(tokens[i]) {
			stark = append(stark, toInt(tokens[i]))
		} else {
			if len(stark) < 2 {
				return 0
			}
			last := stark[len(stark)-1]
			first := stark[len(stark)-2]
			stark = stark[:len(stark)-2]
			stark = append(stark, cac(tokens[i], first, last))
		}
	}
	if len(stark) == 0 {
		return 0
	}
	return stark[0]
}

func symbol(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "/" {
		return true
	}
	return false
}

func toInt(str string) int {
	atoi, _ := strconv.Atoi(str)
	return atoi
}

func cac(str string, first, last int) int {

	str = strings.Trim(str, " ")
	switch str {
	case "+":
		return first + last
	case "-":
		return first - last
	case "*":
		return first * last
	case "/":
		return first / last
	default:
		return 0
	}
}
