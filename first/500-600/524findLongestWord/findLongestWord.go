package main

import (
	"fmt"
	"sort"
)

func main() {
	d := []string{"ale", "apple", "appla", "monkey", "plea"}
	res := findLongestWord("abpcplea", d)
	fmt.Println("res is ", res)
}

/*
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，
返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。
示例 1:
输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]
输出:
"apple"
示例 2:
输入:
s = "abpcplea", d = ["a","b","c"]
输出:
"a"
说明:
    所有输入的字符串只包含小写字母。
    字典的大小不会超过 1000。
    所有输入的字符串长度不会超过 1000。
*/

func findLongestWord(s string, d []string) string {
	sort.Sort(Item(d))

	fmt.Println("d is ", d)
	for _, word := range d {
		if find(word, s) {
			return word
		}
	}
	return ""
}

func find(s, word string) bool {
	if s == word {
		return true
	}
	if len(s) >= len(word) {
		return false
	}

	ll, rl := 0, 0
	for ll < len(s) && rl < len(word) {
		for s[ll] != word[rl] {
			rl++
			if rl >= len(word) {
				return false
			}
		}
		ll++
		rl++
	}
	return ll == len(s)
}

type Item []string

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if len(it[i]) > len(it[j]) {
		return true
	} else if len(it[i]) == len(it[j]) {
		return it[i] < it[j]
	}
	return false
}
func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
