package main

import (
	"fmt"
)

func main() {
	s := "goodgoodwordgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	// words := []string{"foo"}
	sub := findSubstring(s, words)
	fmt.Println("sub is ", sub)
}

/*
给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
示例 1：
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
*/
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
		return res
	}
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}
	ss := []byte(s)
	for i := 0; i <= len(ss)-len(words)*len(words[0]); i++ {
		if find2(ss, i, words, wordMap) {
			res = append(res, i)
		}
	}
	return res
}

func find(ss []byte, start int, words []string) bool {
	// 每次都要构建map，很耗性能
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}
	for i := start; i < start+(len(words[0])*len(words)); i += len(words[0]) {
		word := string(ss[i : i+len(words[0])])
		if wordMap[word] <= 0 {
			return false
		}
		wordMap[word]--
	}
	return true
}

func find2(ss []byte, start int, words []string, wordMap map[string]int) bool {
	// 每次都要构建map，很耗性能
	curMap := make(map[string]int)
	for i := start; i < start+(len(words[0])*len(words)); i += len(words[0]) {

		word := string(ss[i : i+len(words[0])])
		if wordMap[word] <= curMap[word] {
			return false
		}
		curMap[word]++
	}
	return true
}
