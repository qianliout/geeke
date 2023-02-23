package main

import (
	"fmt"
)

func main() {
	fmt.Println('A', 'a')
	fmt.Println('a' - 'A')
	fmt.Println(letterCasePermutation("a12b"))
}

/*
给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。
示例：
输入：S = "a1b2"
输出：["a1b2", "a1B2", "A1b2", "A1B2"]
输入：S = "3z4"
输出：["3z4", "3Z4"]
输入：S = "12345"
输出：["12345"]
提示：
S 的长度不超过12。
S 仅由数字和字母组成。
*/
func letterCasePermutation(S string) []string {
	used := make(map[string]bool)
	res := make([]string, 0)
	queue := make([]string, 0)
	queue = append(queue, S)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if used[first] {
			continue
		}
		used[first] = true
		res = append(res, first)
		queue = append(queue, change(first)...)
	}
	return res
}
func change(s string) []string {
	ss := []byte(s)
	res := make([]string, 0)
	for i := 0; i < len(ss); i++ {
		if ss[i] >= 'a' && ss[i] <= 'z' {
			ss[i] = ss[i] - 32
			res = append(res, string(ss))
			ss[i] = ss[i] + 32
		}
		if ss[i] >= 'A' && ss[i] <= 'Z' {
			ss[i] = ss[i] + 32
			res = append(res, string(ss))
			ss[i] = ss[i] - 32
		}
	}
	return res
}
