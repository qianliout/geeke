package main

import (
	"fmt"
)

func main() {

	/*
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
		["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
	*/
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	words := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	res := usedp(s, words)
	fmt.Println("res is ", res)

}

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/

func wordBreak(s string, wordDict []string) bool {
	mem := make(map[string]bool)
	for _, w := range wordDict {
		mem[w] = true
	}
	path, res, ss := make([]string, 0), false, []byte(s)
	dfs(ss, 0, len(s), path, &res, mem)
	return res
}

// 可以得到答案，但是会超时
func dfs(ss []byte, start, end int, path []string, res *bool, mem map[string]bool) {
	if *res {
		return
	}
	if start >= end {
		if len(path) > 0 {
			*res = true
		}
		return
	}
	for i := start; i <= len(ss); i++ {
		if mem[string(ss[start:i])] {
			path = append(path, string(ss[start:i]))
			dfs(ss, i, end, path, res, mem)
			path = path[:len(path)-1]
		}
	}
}

func usedp(s string, wordDict []string) bool {
	ss := []byte(s)
	mem := make(map[string]bool)
	for _, w := range wordDict {
		mem[w] = true
	}
	// ss[:i]是否可以
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(ss); i++ {
		for j := i - 1; j >= 0; j-- {
			if mem[string(ss[j:i])] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
