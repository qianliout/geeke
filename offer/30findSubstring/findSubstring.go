package main

import (
	"fmt"
)

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	ans := findSubstring(s, words)
	fmt.Println("ans is ", ans)
	m := map[string]int{"3": 3}

	chang(m)
	fmt.Println(m)
}

func chang(m map[string]int) {
	m["3"] = 2

}

/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
示例 1：
输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
注意，单词可能用重复
*/
func findSubstring(s string, words []string) []int {
	ans := make([]int, 0)
	if len(words) == 0 || len(words[0]) == 0 || len(s) < len(words)*len(words[0]) {
		return ans
	}

	left, sing, length, ss := 0, len(words[0]), len(words)*len(words[0]), []byte(s)
	for left+length <= len(s) { // 这里是 <= 也是这个题容易出错的地方
		if eq(ss, left, length, sing, words) {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}

func eq(ss []byte, start, length, sing int, word []string) bool {
	// 因为有重复的单词，且会改变map的值，所以，只能在每个函数的构建
	words := make(map[string]int)
	for _, w := range word {
		words[w]++
	}

	for i := start; i < start+length; i += sing {
		if words[string(ss[i:i+sing])] <= 0 {
			return false
		}
		words[string(ss[i:i+sing])]--
	}
	return true
}
