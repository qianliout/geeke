package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := addStrings("123", "1234")
	fmt.Println("res is ", res)

}

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
提示：
    num1 和num2 的长度都小于 5100
    num1 和num2 都只包含数字 0-9
    num1 和num2 都不包含任何前导零
    你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
*/

func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}
	if len(num2) == 0 {
		return num1
	}
	res := make([]string, 0)
	if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}
	i := 0
	res = append(res, "0")
	for i < len(num1) || i < len(num2) {

		ri, _ := strconv.Atoi(res[len(res)-1])
		n1i := 0
		if i < len(num1) {
			n1i, _ = strconv.Atoi(string(num1[len(num1)-i-1]))
		}
		n2i := 0
		if i < len(num2) {
			n2i, _ = strconv.Atoi(string(num2[len(num2)-i-1]))
		}
		n := ri + n1i + n2i
		if n >= 10 {
			res[len(res)-1] = strconv.Itoa(n % 10)
			res = append(res, strconv.Itoa(n/10))
		} else {
			res[len(res)-1] = strconv.Itoa(n)
			res = append(res, strconv.Itoa(0))
		}
		i++
	}

	ans := ""
	for i := len(res) - 1; i >= 0; i-- {
		if len(ans) == 0 && res[i] == "0" {
			continue
		}
		ans = ans + res[i]
	}
	return ans
}

func addStrings2(num1 string, num2 string) string {
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
	return ans
}
