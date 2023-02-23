package main

import (
	"fmt"
	"strings"
)

func main() {

}

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。
示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"
*/

func multiply(num1 string, num2 string) string {
	ans, n1, n2 := "", []byte(num1), []byte(num2)
	for i := len(num1) - 1; i >= 0; i-- {
		add := 0
		j := len(num2) - 1
		result := ""
		for j >= 0 || add != 0 {
			num := 0
			if j >= 0 {
				num = int(n2[j]-'0') * int(n1[i]-'0')
			}
			num += add
			result = fmt.Sprintf("%d%s", num%10, result)
			add = num / 10
			j--
		}

		ans = addStrings(result+strings.Repeat("0", len(num1)-i-1), ans)
	}
	return ans
}

func addStrings(num1 string, num2 string) string {
	ans, add, n1, n2 := "", 0, []byte(num1), []byte(num2)
	for i, j := len(n1)-1, len(n2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y, res := 0, 0, 0
		if i >= 0 {
			x = int(n1[i] - '0')
		}
		if j >= 0 {
			y = int(n2[j] - '0')
		}
		res = x + y + add
		ans = fmt.Sprintf("%d%s", res%10, ans)
		add = res / 10
	}
	// 去除前面的0
	ansb, i := []byte(ans), 0
	for i < len(ansb) {
		if ansb[i] != '0' {
			break
		}
		i++
	}
	if i == len(ansb) {
		return "0"
	}
	return string(ansb[i:])
}
