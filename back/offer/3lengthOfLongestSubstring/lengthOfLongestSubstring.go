package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := int64(1)
	fmt.Println(unsafe.Sizeof(i)) // 4
	j := 1
	fmt.Println(unsafe.Sizeof(j)) // 4
	u := uint(1)
	fmt.Println(unsafe.Sizeof(u)) // 4

}

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
*/

func lengthOfLongestSubstring(s string) int {
	ss := []byte(s)

	left, right := 0, 0
	ans := 0
	widown := make([]byte, 0)
	for right < len(s) {
		widown = append(widown, ss[right])
		if need(widown) {
			widown = widown[1:]
			left++
		}
		right++
		if len(widown) > ans {
			ans = len(widown)
		}
	}
	return ans
}

func need(widown []byte) bool {
	exit := make(map[byte]bool)
	for _, n := range widown {
		if exit[n] {
			return false
		}
		exit[n] = true
	}
	return true
}

func lengthOfLongestSubstrings(s string) int {
	ss := []byte(s)
	ans, left, right := 0, 0, 0
	exit := make(map[byte]int)
	for right < len(s) {
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
