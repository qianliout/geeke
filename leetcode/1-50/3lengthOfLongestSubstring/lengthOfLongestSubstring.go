package main

import (
	"fmt"
)

func main() {
	a := "abcabcbb"
	substring := lengthOfLongestSubstring(a)
	fmt.Println(substring)
}

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/
func lengthOfLongestSubstring(s string) int {
	ans, ss := 0, []byte(s)
	left, right := 0, 0
	exit := make(map[byte]int)
	for right < len(ss) {
		exit[ss[right]]++
		for exit[ss[right]] > 1 {
			exit[ss[left]]--
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++
	}
	return ans
}

// 这道题的难点是边界的判断
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans, ss := 0, []byte(s)
	left, right := 0, 0
	exit := make(map[byte]int)
	for right < len(ss) {
		exit[ss[right]]++
		for exit[ss[right]] > 1 {
			exit[ss[left]]--
			left++
		}
		if right-left > ans {
			ans = right - left
		}
		right++
	}
	// 这样直接返回是错的，因为空串时会出错,所以必须在前面判空
	return ans + 1
}
