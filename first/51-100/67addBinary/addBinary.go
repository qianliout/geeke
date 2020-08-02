package main

import (
	"fmt"
	"strconv"

	. "outback/leetcode/common"
)

func main() {
	a := "1010"
	b := "1011"
	res := addBinary(a, b)
	fmt.Println("res is ", res)
}

/*
给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字 1 和 0。
示例 1:
输入: a = "11", b = "1"
输出: "100"
示例 2:
输入: a = "1010", b = "1011"
输出: "10101"
提示：
    每个字符串仅由字符 '0' 或 '1' 组成。
    1 <= a.length, b.length <= 10^4
    字符串如果不是 "0" ，就都不含前导零。
*/
func addBinary(a string, b string) string {
	maxLenght := Max(len(a), len(b))
	res := make([]int, maxLenght+1)
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	if len(a) < len(b) {
		b, a = a, b
	}
	as := []byte(a)
	bs := []byte(b)
	minLength := Min(len(a), len(b))
	for i := maxLenght - 1; i >= 0; i-- {
		num := int(as[i] - 48)
		j := i - (maxLenght - minLength)
		if j >= 0 {
			num = num + int(bs[j]-48)
		}
		num += res[i+1]
		if num > 1 {
			res[i+1] = num - 2
			res[i] += 1
		} else {
			res[i+1] = num
		}

	}
	if res[0] == 0 {
		res = res[1:]
	}
	ans := ""
	for _, k := range res {
		ans = ans + strconv.Itoa(k)
	}
	return ans
}
