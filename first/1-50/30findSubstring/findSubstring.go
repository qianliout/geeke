package main

import (
	"fmt"
)

func main() {
	// defer profile.Start().Stop()
	words := []string{"fooo", "barr", "wing", "ding", "wing"}
	// words := []string{"a"}
	res := findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", words)
	// res := findSubstring("a", words)
	fmt.Println("res is ", res)
}

/*
 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
 示例 1：
 输入：
   s = "barfoothefoobarman",
   words = ["foo","bar"]
 输出：[0,9]
 解释：
 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
 输出的顺序不重要, [9,0] 也是有效答案。
 示例 2：
 输入：
   s = "wordgoodgoodgoodbestword",
   words = ["word","good","best","word"]
 输出：[]
*/
// 暴力法，只是这个要超出时间限制
func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(s) == 0 || len(words) == 0 {
		return res
	}
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word] += 1
	}
	n := len(words[0])
	wordNum := len(words)
	byteSlice := []byte(s)
	for i := 0; i < len(s)-wordNum*n+1; i++ {
		cur := i
		// curWordsMap := make(map[string]int) // 性能杀手
		// bytes, _ := json.Marshal(wordsMap)
		// json.Unmarshal(bytes, &curWordsMap)
		curWordsMap := make(map[string]int)
		matched := 0
		for cur < len(byteSlice) {
			word := string(byteSlice[cur : cur+n])
			window, _ := curWordsMap[word]
			need, exit := wordsMap[word]

			if exit && window < need {
				curWordsMap[word] += 1
				cur += n
				matched++
			} else {
				// 如果当前的单词不存在给定的words中，就不往后匹配，
				break
			}
			if matched == wordNum {
				res = append(res, i)
				break
			}
		}
	}
	return res
}
