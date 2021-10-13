package main

import (
	"fmt"
)

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	res := palindromePairs2(words)
	fmt.Println("res is ", res)
}

/*
给定一组 互不相同 的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
示例 1：
输入：["abcd","dcba","lls","s","sssll"]
输出：[[0,1],[1,0],[3,2],[2,4]]
解释：可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]
示例 2：
输入：["bat","tab","cat"]
输出：[[0,1],[1,0]]
解释：可拼接成的回文串为 ["battab","tabbat"]
*/

// 暴力法，会超时
func palindromePairs(words []string) [][]int {
	ans := make([][]int, 0)
	if len(words) <= 1 {
		return ans
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j {
				tem := words[i] + words[j]
				if ispalind([]byte(tem)) {
					ans = append(ans, []int{i, j})
				}
			}
		}
	}
	return ans
}

func palindromePairs2(words []string) [][]int {

	ans := make([][]int, 0)
	if len(words) <= 1 {
		return ans
	}
	wordDict := make(map[string]int)
	for i, w := range words {
		wordDict[w] = i + 1
	}

	for i := 0; i < len(words); i++ {
		ww := []byte(words[i])
		for j := 0; j <= len(ww); j++ {
			tem1 := ww[:j] // 前缀
			tem2 := ww[j:] // 后缀
			// 当word的前缀的反转在字典中，且不是word自身，且word剩下部分是回文(空也是回文)
			// 则说明存在能与word组成回文的字符串
			tem1r := reversal(tem1)
			if ispalind(tem2) && wordDict[tem1r] != i+1 && wordDict[tem1r] > 0 {
				ans = append(ans, []int{i, wordDict[tem1r] - 1})
			}

			if j > 0 {
				// 当word的后缀的反转在字典中，且不是word自身，且word剩下部分是回文(空也是回文)
				// 则说明存在能与word组成回文的字符串
				// 注意：因为是后缀，所以至少要从word的第二位算起，所以j>0
				tem2r := reversal(tem2)
				if ispalind(tem1) && wordDict[tem2r] != i+1 && wordDict[tem2r] > 0 {
					ans = append(ans, []int{wordDict[tem2r] - 1, i})
				}
			}
		}
	}
	return ans
}

func ispalind(word []byte) bool {
	if len(word) <= 1 {
		return true
	}

	left, right := 0, len(word)-1
	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func reversal(word []byte) string {
	l := len(word)
	w := make([]byte, l)
	for i := 0; i < len(word); i++ {
		w[l-i-1] = word[i]
	}
	return string(w)
}
