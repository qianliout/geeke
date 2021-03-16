package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "aaa"
	d := []string{"aaa", "aa", "a"}

	ans := findLongestWord(s, d)
	fmt.Println(ans)
}

/*
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。
如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。
示例 1:
输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]
输出:
"apple"
示例 2:
输入:
s = "abpcplea", d = ["a","b","c"]
输出:
"a"
说明:
所有输入的字符串只包含小写字母。
字典的大小不会超过 1000。
所有输入的字符串长度不会超过 1000。
*/
func findLongestWord(s string, dictionary []string) string {
	ss := []byte(s)
	sort.Sort(Item(dictionary))
	fmt.Println(dictionary)
	for _, d := range dictionary {
		if find(ss, d) {
			return d
		}
	}
	return ""
}

func find(ss []byte, word string) bool {
	words := []byte(word)
	left, start := 0, 0
	for left < len(ss) {
		if ss[left] == words[start] {
			start++
			if start == len(words) {
				return true
			}
		}
		left++
	}
	return start == len(words)
}

type Item []string

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if len(it[i]) > len(it[j]) || (len(it[i]) == len(it[j]) && it[i] < it[j]) {
		return true
	}
	return false
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
