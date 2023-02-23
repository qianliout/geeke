package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	n := lengthOfLongestSubstring3(s)
	fmt.Println("n is ", n)
}

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func lengthOfLongestSubstring(s string) int {
	left, right, res := 0, 0, 0
	if len(s) == 0 {
		return res
	}
	window := make(map[uint8]int)
	for right < len(s) {
		chr := s[right]
		window[chr]++
		right++
		for window[chr] > 1 {
			cha := s[left]
			window[cha]--
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	lastOccrued := make(map[rune]int)
	maxLength, start := 0, 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccrued[ch]
		if lastI >= start && ok {
			start++
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccrued[ch] = i
	}
	return maxLength
}

func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	ss := []byte(s)
	left, right := 0, 0
	widow := make(map[byte]int)
	ans := 0
	for right < len(ss) {
		// 扩大窗口
		cha := ss[right]
		widow[cha]++
		right++

		// 缩小窗口
		for widow[cha] > 1 {
			widow[ss[left]]--
			left++
		}
		// 获取结果
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}
