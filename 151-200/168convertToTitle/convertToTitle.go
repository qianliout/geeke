package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(byte(25 + 64)))
	res := convertToTitle(52)
	fmt.Println("res is ", res)
}

/*
给定一个正整数，返回它在 Excel 表中相对应的列名称。
例如，
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
示例 1:
输入: 1
输出: "A"
示例 2:
输入: 28
输出: "AB"
示例 3:
输入: 701
输出: "ZY"
*/
func convertToTitle(n int) string {
	if n == 0 {
		return ""
	}
	ans := ""
	for n > 26 {
		num := n % 26
		if num == 0 {
			ans = "Z" + ans
			n = (n / 26) - 1
		} else {
			ans = string(byte(num+64)) + ans
			n = n / 26
		}

	}
	ans = string(byte(n+64)) + ans
	return ans
}
