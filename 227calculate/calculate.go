package main

import (
	"strconv"
	"strings"
)

func main() {

}

/*
基本计算器 II

1 <= s.length <= 3 * 105
s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
s 表示一个 有效表达式
*/
// 没有括号，就要简单的多
func calculate(s string) int {
	num, sign, numStark := 0, '+', make([]int, 0)
	ss := []byte(strings.ReplaceAll(s, " ", ""))
	for i := range ss {
		if ss[i] >= '0' && ss[i] <= '9' {
			n, _ := strconv.Atoi(string(ss[i]))
			num = num*10 + n
			// 如果最后一个是数字，也是要计算的
			if i != len(ss)-1 {
				continue
			}
		}
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
		num, sign = 0, int32(ss[i])
	}
	sum := 0
	for _, n := range numStark {
		sum += n
	}
	return sum
}
