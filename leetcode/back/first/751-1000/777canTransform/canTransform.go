package main

import (
	"fmt"
)

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
	fmt.Println(canTransform("XXRXXXXRXXXXXXRXXXXLXXXXLXLXXXXXXRXXXXLX", "XXXXXXXXXXXXRRRLLLXXXXXXXXXXXXXXXXXXRLXX"))
}

/*
在一个由 'L' , 'R' 和 'X' 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。现给定起始字符串start和结束字符串end，请编写代码，当且仅当存在一系列移动操作使得start可以转换成end时， 返回True。
示例 :
输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
输出: True
解释:
我们可以通过以下几步将start转换成end:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX
提示：
1 <= len(start) = len(end) <= 10000。
start和end中的字符串仅限于'L', 'R'和'X'。
*/
// 这种方式可以得到结果，但是会超时
func canTransform(start string, end string) bool {
	used := make(map[string]bool)
	queue := make([]string, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if used[first] {
			continue
		}
		if first == end {
			return true
		}
		used[first] = true
		queue = append(queue, change(first)...)
	}
	return false
}

func change(start string) []string {
	res := make([]string, 0)
	ss := []byte(start)
	for i := 0; i < len(start)-1; i++ {
		if ss[i] == 'R' && ss[i+1] == 'X' {
			ss[i] = 'X'
			ss[i+1] = 'R'
			res = append(res, string(ss))
			ss[i] = 'R'
			ss[i+1] = 'X'
		}
		if ss[i] == 'X' && ss[i+1] == 'L' {
			ss[i] = 'L'
			ss[i+1] = 'X'
			res = append(res, string(ss))
			ss[i] = 'X'
			ss[i+1] = 'L'
		}
	}
	return res
}
