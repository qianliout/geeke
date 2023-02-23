package main

import (
	"fmt"
)

func main() {
	window := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println("widnown is ", window)
}

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow(s string, t string) string {
	left, right, ans := 0, 0, ""
	ss, tt := []byte(s), []byte(t)

	ttMap := make(map[byte]int)
	for i := range tt {
		ttMap[tt[i]]++
	}
	ssMap := make(map[byte]int)

	for right < len(s) {
		ssMap[ss[right]]++

		for has(ssMap, ttMap) {
			now := string(ss[left : right+1])
			if ans == "" || (ans != "" && len(now) < len(ans)) {
				ans = now
			}
			ssMap[ss[left]]--
			left++
		}
		right++
	}
	return ans
}

func has(ss1map, ss2map map[byte]int) bool {
	for k, v := range ss2map {
		if ss1map[k] < v {
			return false
		}
	}
	return true
}
