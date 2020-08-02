package main

import (
	"fmt"
)

func main() {
	s := "aab"
	res := lengthOfLongestSubstring(s)
	fmt.Println("res is ", res)
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
	if len(s) == 0 {
		return 0
	}
	res := 0
	left, right := 0, 0
	exit := make(map[byte]int)
	for right < len(s) {
		exit[s[right]]++
		right++
		if exit[s[right-1]] > 1 {
			for left < right {
				exit[s[left]]--
				left++
			}
		}
		if right-left > res {
			res = right - left
		}
	}
	if right-left > res {
		res = right - left
	}
	return res
}
