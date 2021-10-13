package main

import (
	"fmt"
)

func main() {
	res := removeInvalidParentheses("((()((s((((()")
	fmt.Println("res is ", res)
}

/*
删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
说明: 输入可能包含了除 ( 和 ) 以外的字符。
示例 1:
输入: "()())()"
输出: ["()()()", "(())()"]
示例 2:
输入: "(a)())()"
输出: ["(a)()()", "(a())()"]
示例 3:
输入: ")("
输出: [""]
*/

// 题目要求删除最少的括号，所以，使用bfs是最好了，有结果就返回
// 使用bfs，先删除一个括号，进行测试如果有一个是合法的就返回，如果没有，就把上一轮已删除一个括号的队列再删除一下括号
func removeInvalidParentheses(s string) []string {
	queue := make([][]byte, 0)
	queue = append(queue, []byte(s))
	for len(queue) > 0 {
		res := make([]string, 0)
		for _, i := range queue {
			if isValid(i) {
				res = append(res, string(i))
			}
		}
		if len(res) > 0 {
			return duplicate(res)
		}
		// 这一层没有，就到下一层
		nextLevel := make([][]byte, 0)
		for _, q := range queue {
			for i := 0; i < len(q); i++ {
				if q[i] == '(' || q[i] == ')' {
					// 这样写gc压力大，测试不能通过
					nextLevel = append(nextLevel, append(append([]byte{}, q[:i]...), q[i+1:]...))
				}
			}
		}
		queue = nextLevel
	}
	return []string{}
}

func isValid(ss []byte) bool {
	cnt := 0
	for _, i := range ss {
		if i == '(' {
			cnt++
		} else if i == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

func duplicate(ss []string) []string {
	res := make([]string, 0)
	d := make(map[string]bool)
	for _, s := range ss {
		if !d[s] {
			res = append(res, s)
			d[s] = true
		}
	}
	return res
}
