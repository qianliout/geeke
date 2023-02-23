package main

import (
	"fmt"
	"strings"
)

func main() {
	res := wordPattern("abab", "dog cat cat dog")
	fmt.Println(res)
}

/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
示例1:
输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:
输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:
输入: pattern = "abba", str = "dog dog dog dog"
输出: false
*/

func wordPattern(pattern string, str string) bool {
	ss := []byte(pattern)

	strs := strings.Split(str, " ")
	if len(ss) != len(strs) {
		return false
	}

	pMap := make(map[byte]int)
	sMap := make(map[string]int)

	for i := 0; i < len(ss); i++ {
		pMap[ss[i]] = i
		sMap[strs[i]] = i
	}

	// 判断
	for i := 0; i < len(ss); i++ {
		if pMap[ss[i]] != sMap[strs[i]] {
			return false
		}
	}
	return true
}
