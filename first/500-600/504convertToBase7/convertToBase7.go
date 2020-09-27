package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := convertToBase7(100)
	fmt.Println("res is ", res)
}

/*
给定一个整数，将其转化为7进制，并以字符串形式输出。
示例 1:
输入: 100
输出: "202"
示例 2:
输入: -7
输出: "-10"
注意: 输入范围是 [-1e7, 1e7] 。
*/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	preNum := num

	ans := ""
	if num < 0 {
		num = -num
	}

	for num > 0 {
		ans = strconv.Itoa(num%7) + ans
		num /= 7
	}
	if preNum <= 0 {
		return "-" + ans
	}
	return ans
}

func convertToBase72(num int) string {
	if num == 0 {
		return "0"
	} else if num < 0 {
		return "-" + convertToBase7(-num)
	} else if num < 7 {
		return strconv.Itoa(num)
	}
	return convertToBase7(num/7) + strconv.Itoa(num%7)
}
