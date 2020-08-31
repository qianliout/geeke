package main

import (
	"fmt"
)

func main() {
	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	res := wordBreak(s, wordDict)
	fmt.Println("res is ", res)
	
}

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，
使得句子中所有的单词都在词典中。返回所有这些可能的句子。
说明：
    分隔时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。
示例 1：
输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
示例 2：
输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：
输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
*/

// 这种写法没有错，只是在特殊情况下（都是重复的字母），时间会很长
// s ="aaaaaaaaaaaaaaaaaaa"
// worddict=["aa","a","aaa"]
func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]interface{})
	for _, w := range wordDict {
		wordMap[w] = nil
	}
	// dp[i]表示i是否是分隔点
	dp := make(map[int]bool)
	dp[0] = true
	
	ss := []byte(s)
	
	// strs[i]表示s[0:i]这个字符串分隔的结果
	strs := make([][]string, len(s)+1)
	// 初始化
	// for i := 0; i < len(strs); i++ {
	// 	strs[i] = make([]string, 0)
	// }
	
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			str := string(ss[j:i])
			_, ok := wordMap[str]
			if ok && dp[j] {
				dp[i] = true
				if j == 0 {
					strs[i] = append(strs[i], str)
				} else {
					for _, st := range strs[j] {
						sb := st + " " + str
						strs[i] = append(strs[i], sb)
					}
				}
			}
		}
	}
	return strs[len(s)]
}

// 进行预处理
func wordBreaks(s string, wordDict []string) []string {
	wordMap := make(map[string]interface{})
	for _, w := range wordDict {
		wordMap[w] = nil
	}
	// dp[i]表示i是否是分隔点
	dp := make(map[int]bool)
	dp[0] = true
	
	ss := []byte(s)
	
	// strs[i]表示s[0:i]这个字符串分隔的结果
	strs := make([][]string, len(s)+1)
	// 初始化
	for i := 0; i < len(strs); i++ {
		strs[i] = make([]string, 0)
	}
	// 预处理
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			str := string(ss[j:i])
			if _, ok := wordMap[str]; dp[j] && ok {
				dp[i] = true
			}
		}
	}
	if !dp[len(s)] {
		return []string{}
	}
	
	for i := 1; i <= len(s); i++ {
		if !dp[i] {
			continue
		}
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			str := string(ss[j:i])
			_, ok := wordMap[str]
			if ok && dp[j] {
				dp[i] = true
				if j == 0 {
					strs[i] = append(strs[i], str)
				} else {
					for _, st := range strs[j] {
						sb := st + " " + str
						strs[i] = append(strs[i], sb)
					}
				}
			}
		}
	}
	return strs[len(s)]
}
