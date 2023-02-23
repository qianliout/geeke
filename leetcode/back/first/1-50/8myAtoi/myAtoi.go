package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	str := "   +00012words "

	fmt.Println('+')
	res := myAtoi(str)
	fmt.Println("res is ", res)
}

/*
你来实现一个 atoi 函数，使其能将字符串转换成整数。
首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。
当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。
该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。
注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
在任何情况下，若函数不能进行有效的转换时，请返回 0。
说明：
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
示例 1:
输入: "42"
输出: 42
示例 2:
输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。
     我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。
示例 3:
输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:
输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。
     因此无法执行有效的转换。
示例 5:
输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。
     因此返回 INT_MIN (−231) 。
*/

// 这个方法是找到每一个数据，我题目中要求的是第一个就要是数字
func myAtoi(str string) int {
	numberMap := make(map[byte]int)
	for i := 48; i < 58; i++ {
		numberMap[byte(i)] = i - 48
	}
	negative := false
	begen := false

	num := 0

	str = strings.Trim(str, " ")
	byteSliece := []byte(str)
	if len(byteSliece) <= 0 {
		return 0
	}
	fmt.Println(byteSliece[0])
	fuck := false
	if byteSliece[0] != byte(45) && numberMap[byteSliece[0]] > 0 {
		fuck = true
	} else if byteSliece[0] == byte(45) && len(byteSliece) > 1 && numberMap[byteSliece[1]] > 0 {
		fuck = true
		negative = true
	} else if byteSliece[0] == byte(43) && len(byteSliece) > 1 && numberMap[byteSliece[1]] > 0 {
		byteSliece = byteSliece[1:]
	}

	if fuck == true {
		for _, va := range byteSliece {
			n, exit := numberMap[va]
			if exit {
				// 先判断溢出
				if !negative && ((num > math.MaxInt32/10) || (num == math.MaxInt32 && n > 7)) {
					return math.MaxInt32
				} else if negative && ((-num < math.MinInt32/10) || (-num == math.MinInt32 && n < -8)) {
					return math.MinInt32
				}
				num = num*10 + n
				begen = true
			} else if begen && !exit {
				break
			}
		}
	}
	if negative {
		num = -num
	}
	return num
}
