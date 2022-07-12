package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("AAA"))
}

/*
给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回 该列名称对应的列序号 。
例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
*/
func titleToNumber(columnTitle string) int {
	ans, ss := 0, []byte(columnTitle)

	for i := range ss {
		ans = ans*26 + (int(ss[i]) - 65 + 1)
	}
	return ans
}
