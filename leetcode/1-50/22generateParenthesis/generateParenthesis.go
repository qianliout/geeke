package main

import (
	"fmt"
)

func main() {
	res := generateParenthesis2(3)
	fmt.Println("res is ", res)
}

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
*/
func generateParenthesis2(n int) []string {
	path := make([]byte, 0)
	res := make([]string, 0)
	backtrack(path, n, 0, 0, &res)
	return res
}

func backtrack(path []byte, n, left, right int, res *[]string) {
	// left,right表示已经使用的括号数量,
	// result 表示当前已经生成的结果，
	if left == right && right == n {
		*res = append(*res, string(append([]byte{}, path...)))
		return
	}
	if left >= right && left < n {
		path = append(path, '(')
		backtrack(path, n, left+1, right, res)
		path = path[:len(path)-1]
	}
	if right < n {
		path = append(path, ')')
		backtrack(path, n, left, right+1, res)
		path = path[:len(path)-1]
	}
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		// 递归终止条件，因为下面是:res := "(" + left + ")" + right,所以这里必须是返回[]string{""}
		return []string{""}
		// 这样写是错的
		// return []string{}
	}
	result := make([]string, 0)
	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis(i) {
			for _, right := range generateParenthesis(n - i - 1) {
				// 两种写法都是对的
				// res := "(" + right + ")" + left
				res := "(" + left + ")" + right
				result = append(result, res)
			}
		}
	}
	return result
}
