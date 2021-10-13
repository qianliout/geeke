package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := calculate3("(1+3-(40+543+2)-3)+(6+83)+20")
	fmt.Println("res is ", res)
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

type item struct {
	value  int
	symbol string
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.Replace(s, " ", "", -1)
	ss := []byte(s)
	length := len(ss)
	symbolMap := map[string]bool{"(": true, ")": true, "+": true, "-": true}
	i := 0
	itemStark := make([]item, 0)

	for i < length {
		if symbolMap[string(ss[i])] {
			itemStark = append(itemStark, item{symbol: string(ss[i])})
		} else {
			num, _ := strconv.Atoi(string(ss[i]))
			for i+1 < length && !symbolMap[string(ss[i+1])] {
				r, _ := strconv.Atoi(string(ss[i+1]))
				num = num*10 + r
				i++
			}
			itemStark = append(itemStark, item{value: num})
		}
		i++
	}
	numStark := make([]int, 0)
	symbolStak := make([]string, 0)

	for _, it := range itemStark {
		if len(it.symbol) > 0 && it.symbol != ")" {
			symbolStak = append(symbolStak, it.symbol)
		}
		if len(it.symbol) > 0 && it.symbol == ")" {
			if len(symbolStak) < 1 || symbolStak[len(symbolStak)-1] != "(" {
				fmt.Println("非法1")
				return 0
			}
			symbolStak = symbolStak[:len(symbolStak)-1]
			if len(symbolStak) > 0 && symbolStak[len(symbolStak)-1] != "(" {
				if len(numStark) < 2 {
					fmt.Println("非法2", numStark, symbolStak)
					return 0
				}
				var newNum int
				lastSymbol := symbolStak[len(symbolStak)-1]
				symbolStak = symbolStak[:len(symbolStak)-1]

				if lastSymbol == "-" {
					newNum = numStark[len(numStark)-2] - numStark[len(numStark)-1]
				} else if lastSymbol == "+" {
					newNum = numStark[len(numStark)-2] + numStark[len(numStark)-1]
				}
				numStark = numStark[:len(numStark)-2]
				numStark = append(numStark, newNum)
			}

		}

		// 说明是数字
		if len(it.symbol) == 0 {
			if len(symbolStak) == 0 || symbolStak[len(symbolStak)-1] == "(" || len(numStark) == 0 {
				numStark = append(numStark, it.value)
			} else {
				for {
					lastSymbol := symbolStak[len(symbolStak)-1]
					lastNum := numStark[len(numStark)-1]
					numStark = numStark[:len(numStark)-1]

					if lastSymbol == "-" {
						lastNum = lastNum - it.value
					} else if lastSymbol == "+" {
						lastNum = lastNum + it.value
					}
					numStark = append(numStark, lastNum)
					symbolStak = symbolStak[:len(symbolStak)-1]
					if len(symbolStak) == 0 || symbolStak[len(symbolStak)-1] == "(" {
						break
					}
				}
			}
		}
	}
	return numStark[0]
}

func calculate2(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	res := 0
	operator := 1
	stark := make([]int, 0)
	num := 0
	if len(s) == 0 {
		return res
	}
	for _, ch := range []byte(s) {
		if string(ch) == "(" {
			stark = append(stark, res)
			stark = append(stark, operator)
			res = 0
			operator = 1
			num = 0
		} else if string(ch) == ")" {
			res = res + num*operator
			res = res * stark[len(stark)-1]
			res = res + stark[len(stark)-2]
			stark = stark[:len(stark)-2]
			num = 0
			operator = 1
		} else if string(ch) == "+" {
			res = res + num*operator
			operator = 1
			num = 0
		} else if string(ch) == "-" {
			res = res + num*operator
			operator = -1
			num = 0
		} else {
			n, _ := strconv.Atoi(string(ch))
			num = num*10 + n
		}
	}
	// 要注意的是，最后可能是数字，所以就会存在漏加的情况，1+1 就可以验证
	return res + operator*num
}


// 这个方法一定要滚瓜烂熟
func calculate3(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if len(s) == 0 {
		return 0
	}
	res, num, operator := 0, 0, 1
	numStack := make([]int, 0)
	ss := []byte(s)
	for _, ch := range ss {
		if ch == '(' {
			res = res + num*operator
			numStack = append(numStack, res)
			numStack = append(numStack, operator)
			operator = 1
			res = 0
			num = 0
		} else if ch == ')' {
			res += num * operator
			res = res * numStack[len(numStack)-1]
			res += numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-2]
			num = 0
			operator = 1
		} else if ch == '+' {
			res += num * operator
			operator = 1
			num = 0

		} else if ch == '-' {
			res += num * operator
			operator = -1
			num = 0
		} else {
			n, _ := strconv.Atoi(string(ch))
			num = num*10 + n
		}
	}
	return res + num*operator
}
