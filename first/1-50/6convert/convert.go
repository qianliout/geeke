package main

import (
	"fmt"
)

func main() {
	ans := convert("LEETCODEISHIRING", 3)
	fmt.Println("ans is ", ans)
}

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);
示例 1:
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:
L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows <= 1 {
		return s
	}

	index, dir := 0, 1
	res := make([][]byte, numRows)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if res[index] == nil {
			res[index] = make([]byte, 0)
		}
		res[index] = append(res[index], ss[i])
		index += dir
		if index == numRows-1 || index == 0 {
			dir = -dir
		}
	}
	// 组装结果
	ans := make([]byte, 0)
	for i := 0; i < len(res); i++ {
		ans = append(ans, res[i]...)
	}

	return string(ans)
}
