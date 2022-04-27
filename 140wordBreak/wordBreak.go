package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	res := wordBreak(s, wordDict)
	fmt.Println("res is ", res, len(res))
}

func wordBreak(s string, wordDict []string) []string {
	mem := make(map[string]bool)
	for _, n := range wordDict {
		mem[n] = true
	}
	path, res := make([]string, 0), make([]string, 0)
	dfs([]byte(s), 0, path, &res, mem)
	return res
}

func dfs(ss []byte, start int, path []string, res *[]string, mem map[string]bool) {
	if start >= len(ss) {
		*res = append(*res, strings.Join(path, " "))
		return
	}
	for i := start; i < len(ss); i++ {
		if mem[string(ss[start:i+1])] {
			path = append(path, string(ss[start:i+1]))
			dfs(ss, i+1, path, res, mem)
			path = path[:len(path)-1]
		}
	}
}
