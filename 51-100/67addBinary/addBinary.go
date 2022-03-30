package main

import (
	"fmt"
)

func main() {
	a := "0"
	b := "1"
	binary := addBinary(a, b)
	fmt.Println("binary is ", binary)
}

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 1 和 0。
示例 1:
输入: a = "11", b = "1"
输出: "100"
示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
*/
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ans := make([]int, len(a)+1)

	aa, bb := []byte(a), []byte(b)
	last, l := 0, len(a)-len(b)
	for i := len(b) - 1; i >= 0; i-- {
		num := int(aa[i+l]-'0'+bb[i]-'0') + last
		if num >= 2 {
			last = 1
			num = num - 2
		} else {
			last = 0
		}
		ans[i+l+1] = num
	}

	for i := l - 1; i >= 0; i-- {
		num := int(aa[i]-'0') + last
		if num >= 2 {
			last = 1
			num = num - 2
		} else {
			last = 0
		}
		ans[i+1] = num
	}
	if last > 0 {
		ans[0] = last
	} else {
		ans = ans[1:]
	}
	res := ""
	for i := range ans {
		res = fmt.Sprintf("%s%d", res, ans[i])
	}
	return res
}
