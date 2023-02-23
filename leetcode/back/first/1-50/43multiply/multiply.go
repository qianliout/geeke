package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := addStrings("00", "00")
	fmt.Println(res)
}

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"
说明：
    num1 和 num2 的长度小于110。
    num1 和 num2 只包含数字 0-9。
    num1 和 num2 均不以零开头，除非是数字 0 本身。
    不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

func multiply(num1 string, num2 string) string {
	ans := ""
	for i := len(num2) - 1; i >= 0; i-- {
		add := 0
		j := len(num1) - 1
		result := ""
		for j >= 0 || add != 0 {
			num := 0
			if j >= 0 {
				num = int(num2[i]-'0') * int(num1[j]-'0')
			}
			num += add
			result = strconv.Itoa(num%10) + result
			add = num / 10
			j--
		}

		ans = addStrings(result+strings.Repeat("0", len(num2)-i-1), ans)
	}
	return ans
}
func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		result := add + x + y
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	// 去除前面的0
	i := 0
	for i < len(ans) {
		if ans[i] == '0' {
			i++
		} else {
			break
		}
	}
	if i > 0 {
		ans = string([]byte(ans)[i:])
		if len(ans) == 0 {
			ans = "0"
		}
		return ans
	}
	return ans
}
