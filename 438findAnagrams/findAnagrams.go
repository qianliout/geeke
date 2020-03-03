package main

import (
	"fmt"
)

/*
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
说明：
    字母异位词指字母相同，但排列不同的字符串。
    不考虑答案输出的顺序。
示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 示例 2:
输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println("res is ", res)
}

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(p) == 0 {
		return res
	}
	needs := make(map[uint8]int)
	window := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		chr := p[i]
		needs[chr]++
	}

	left, right, matched := 0, 0, 0
	for right < len(s) {
		cha := s[right]

		_, exit := needs[cha]
		right++

		if exit {
			window[cha]++
			if needs[cha] == window[cha] {
				matched++
			}

			for matched == len(needs) {
				fmt.Println("left is ", left, "right is ", right)
				if right-left == len(p) {
					res = append(res, left)
				}

				chr := s[left]
				if _, exit := needs[chr]; exit {
					window[chr]--
					if window[chr] < needs[chr] {
						matched--
					}
				}
				left++
			}
		}
		if !exit {
			window = make(map[uint8]int)
			left = right
			matched = 0
		}
	}
	return res
}
