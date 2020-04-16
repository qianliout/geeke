package main

import (
	"fmt"
)

func main() {
	fmt.Println('A', 'Z')
	res := titleToNumber("AZ")
	fmt.Println("res is ", res)
}

/*
给定一个Excel表格中的列名称，返回其相应的列序号。
例如，
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...
示例 1:
输入: "A"
输出: 1
示例 2:
输入: "AB"
输出: 28
示例 3:
输入: "ZY"
输出: 701
*/

func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[string]int)
	for i := 65; i <= 90; i++ {
		m[string(byte(i))] = i - 64
	}
	num := 0
	ss := []byte(s)
	for i := 0; i < len(s); i++ {
		r := m[string(ss[i])]
		num = num*26 + r
	}
	return num
}
