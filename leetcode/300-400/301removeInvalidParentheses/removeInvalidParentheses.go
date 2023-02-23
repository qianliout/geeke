package main

import (
	"fmt"
)

func main() {
	s := "(a)())()"
	parentheses := removeInvalidParentheses(s)
	fmt.Println("parentheses ", parentheses)
}

/*
给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
返回所有可能的结果。答案可以按 任意顺序 返回。
示例 1：
输入：s = "()())()"
输出：["(())()","()()()"]
*/
func removeInvalidParentheses2(s string) []string {
	ss := []byte(s)
	ans := make([]string, 0)
	exit := make(map[string]bool)
	dfs(ss, 0, &ans, exit)
	// 查最大的
	res, length := make([]string, 0), 0
	for i := range ans {
		if len(ans[i]) > length {
			length = len(ans[i])
		}
	}

	for i := range ans {
		if len(ans[i]) == length {
			res = append(res, ans[i])
		}
	}

	return res
}

func removeInvalidParentheses(s string) []string {
	return bfs([]byte(s))
}

func dfs(ss []byte, start int, res *[]string, exit map[string]bool) {
	if valid(ss) {
		st := toString(ss)
		if !exit[st] {
			*res = append(*res, toString(ss))
			exit[st] = true
		}
	}

	for i := start; i < len(ss); i++ {
		if ss[i] != '(' && ss[i] != ')' && ss[i] != '@' {
			continue
		}
		pre := ss[i]
		ss[i] = '@'
		dfs(ss, i+1, res, exit)
		ss[i] = pre
	}
}

func bfs(ss []byte) []string {
	fond, queue, ans, exit := false, make([][]byte, 0), make([]string, 0), make(map[string]bool)
	queue = append(queue, ss)
	exit[toString(ss)] = true
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if valid(first) {
			ans = append(ans, toString(first))
			fond = true
		}
		if fond {
			continue
		}

		for i := 0; i < len(first); i++ {
			if first[i] != '(' && first[i] != ')' {
				continue
			}
			tmp := append([]byte{}, first...)
			tmp[i] = '@'
			if !exit[toString(tmp)] {
				queue = append(queue, tmp)
				exit[toString(tmp)] = true
			}
		}
	}
	return ans
}

func toString(ss []byte) string {
	ans := make([]byte, 0)
	for i := range ss {
		if ss[i] != '@' {
			ans = append(ans, ss[i])
		}
	}
	return string(ans)
}

func valid(ss []byte) bool {
	stark := make([]byte, 0)

	for i := range ss {
		if ss[i] == '(' {
			stark = append(stark, '(')
		} else if ss[i] == ')' {
			if len(stark) == 0 {
				return false
			}
			stark = stark[:len(stark)-1]
		}
	}
	return len(stark) == 0
}
