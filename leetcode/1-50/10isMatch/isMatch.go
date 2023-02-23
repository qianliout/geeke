package main

import "fmt"

func main() {
	fmt.Print("hello")
}

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
示例 1：
输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
*/
func isMatch(s string, p string) bool {
	ss, pp := []byte(s), []byte(p)
	res := true
	return match(ss, pp, &res)
}

func match(ss, pp []byte, res *bool) bool {
	if len(pp) == 0 {
		return len(ss) == 0
	}
	if !(*res) {
		return false
	}
	first := len(ss) > 0 && (ss[0] == pp[0] || pp[0] == '.')
	if len(pp) >= 2 && pp[1] == '*' {
		return match(ss, pp[2:], res) || (first && match(ss[1:], pp, res))
	}
	return first && match(ss[1:], pp[1:], res)
}

func isMatch2(s string, p string) bool {
	ss, pp := []byte(s), []byte(p)
	return match2(ss, pp)
}

func match2(ss, pp []byte) bool {
	if len(pp) == 0 {
		return len(ss) == 0
	}
	first := len(ss) > 0 && (ss[0] == pp[0] || pp[0] == '.')
	if len(pp) >= 2 && pp[1] == '*' {
		return match2(ss, pp[2:]) || (first && match2(ss[1:], pp))
	}
	return first && match2(ss[1:], pp[1:])
}
