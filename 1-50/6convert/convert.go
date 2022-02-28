package main

import (
	"fmt"
)

func main() {
	ans := convert("PAYPALISHIRING", 3)
	fmt.Println(ans)

}

/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：
string convert(string s, int numRows);
示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"
*/
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows <= 1 {
		return s
	}
	res := make(map[int][]byte)
	ss, index, dir := []byte(s), 0, 1

	for i := 0; i < len(ss); i++ {
		res[index] = append(res[index], ss[i])
		index += dir
		if index == 0 || index == numRows-1 {
			dir = -dir
		}
	}
	ans := make([]byte, 0)
	for i := 0; i < numRows; i++ {
		ans = append(ans, res[i]...)
	}
	return string(ans)
}
