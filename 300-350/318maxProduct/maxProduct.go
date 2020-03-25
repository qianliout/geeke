package main

import (
	"fmt"
)

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	res := maxProduct(words)
	fmt.Println("res is ", res)
}

/*
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，
并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。
示例 1:
输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "xtfn"。
示例 2:
输入: ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4
解释: 这两个单词为 "ab", "cd"。
示例 3:
输入: ["a","aa","aaa","aaaa"]
输出: 0
解释: 不存在这样的两个单词。
链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
*/

// 这道题首先想到的就是dp
func maxProduct(words []string) int {
	if len(words) <= 1 {
		return 0
	}
	res := 0
	for i := 1; i < len(words); i++ {
		wordI := words[i]
		for j := i - 1; j >= 0; j-- {
			wordJ := words[j]
			if !diff(wordI, wordJ) {
				r := len(wordI) * len(wordJ)
				if r > res {
					res = r
				}
			}
		}
	}
	return res
}

// 判断是否有公共字母
// 如果有重合，则返回true
//最容易想到的办法是hash表，这里使用位操作,
func diff1(word1, word2 string) bool {
	for _, i := range word1 {
		for _, j := range word2 {
			diff := int(i) ^ int(j)
			if diff == 0 {
				return true
			}
		}
	}
	return false
}

// 使用位移操作
func diff(word1, word2 string) bool {
	b1, b2 := 0, 0
	for _, w := range word1 {
		b1 = b1 | (1 << (w - 'a'))
	}
	for _, w := range word2 {
		b2 = b2 | (1 << (w - 'a'))
	}
	return !(b1&b2 == 0)
}
