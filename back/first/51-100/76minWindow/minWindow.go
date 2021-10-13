package main

import (
	"fmt"
	"math"
)

func main() {
	s := "aa"
	t := "aa"
	res := minWindow2(s, t)
	fmt.Println("window is ", res)
}

/*
 76. 最小覆盖子串
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
示例：
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：
    如果 S 中不存这样的子串，则返回空字符串 ""。
    如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	left, right, matched, start, minLen := 0, 0, 0, 0, math.MaxInt32

	window := make(map[uint8]int)
	needs := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	for right < len(s) {
		cha := s[right]

		if _, exit := needs[cha]; exit {
			window[cha]++
			if window[cha] == needs[cha] {
				matched++
			}
		}

		right++
		// 如果全匹配了
		for matched == len(needs) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			cha2 := s[left]
			if _, exit := needs[cha2]; exit {
				window[cha2]--
				if window[cha2] < needs[cha2] {
					matched--
				}
			}
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return string(s[start : start+minLen])
}

func isAll(smap, tmap map[uint8]int) bool {
	for k, v := range smap {
		if tmap[k] != v {
			return false
		}
	}
	return true
}

func minWindow2(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}
	ss, tt := []byte(s), []byte(t)
	left, right, ansLength, ans := 0, 0, math.MaxInt64, ""
	needs := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[tt[i]]++
	}

	smap := make(map[byte]int)
	for right < len(s) {
		// 扩大窗口
		hasAns := false // 表示是否有答案了
		chr := ss[right]
		smap[chr]++
		right++

		// 缩小窗口了
		for hasAll(smap, needs) && left < right {
			hasAns = true
			smap[ss[left]]--
			left++
		}
		// 更新答案了
		if hasAns {
			thisAns := string(ss[left-1 : right])
			if len(thisAns) < ansLength {
				ans = thisAns
				ansLength = len(ans)
			}
		}
	}

	return ans
}

func hasAll(smap, tt map[byte]int) bool {
	for k, v := range tt {
		if n, ok := smap[k]; !ok || n < v {
			return false
		}
	}
	return true
}
