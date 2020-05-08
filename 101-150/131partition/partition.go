package main

import (
	"fmt"
)

func main() {
	s := "efe"
	res := partition(s)
	fmt.Println("res is ", res)
}

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。
示例:
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
*/
var mem map[string]bool

func partition(s string) [][]string {
	res := make([][]string, 0)
	if len(s) == 0 {
		return res
	}
	path := make([]string, 0)
	mem = make(map[string]bool)
	backtracking([]byte(s), 0, len(s), path, &res)
	return res
}

func backtracking(s []byte, start, end int, path []string, res *[][]string) {
	if start == end { // 这里才是理解的难点
		*res = append(*res, append([]string{}, path...))
		return
	}

	for i := start; i < end; i++ {
		// 因为截取字符串是消耗性能的，因此，采用传子串索引的方式判断一个子串是否是回文子串
		// 不是的话，剪枝
		if !checkPalindrome(s, start, i) {
			continue
		}
		path = append(path, string(s[start:i+1]))
		backtracking(s, i+1, end, path, res)
		path = path[:len(path)-1]
	}
}

func checkPalindrome(s []byte, start, end int) bool {
	ss := string(s[start : end+1])
	if is, ok := mem[ss]; ok {
		return is
	}

	if start == end {
		mem[string(s[start])] = true
		return true
	}
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	mem[string(s[start:end+1])] = true
	return true
}
