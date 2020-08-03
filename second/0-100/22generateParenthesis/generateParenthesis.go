package main

import (
	"fmt"
)

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}
	dfs(n, 0, 0, []byte{}, &res)
	return res
}

func dfs(n, left, right int, path []byte, res *[]string) {
	if len(path) == 2*n {
		*res = append(*res, string(path))
		return
	}
	if left < n {
		path = append(path, '(')
		dfs(n, left+1, right, path, res)
		path = path[:len(path)-1]
	}
	if right < left {
		path = append(path, ')')
		dfs(n, left, right+1, path, res)
		path = path[:len(path)-1]
	}
}
