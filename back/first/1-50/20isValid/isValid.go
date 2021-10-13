package main

import (
	"fmt"
)

func main() {
	s := ")"
	res := isValid(s)
	fmt.Println("res is ", res)
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
示例 1:
输入: "()"
输出: true
示例 2:
输入: "()[]{}"
输出: true
示例 3:
输入: "(]"
输出: false
示例 4:
输入: "([)]"
输出: false
示例 5:
输入: "{[]}"
输出: true
链接：https://leetcode-cn.com/problems/valid-parentheses
*/
func isValid(s string) bool {
	mapping := map[string]string{")": "(", "}": "{", "]": "["}

	//if len(s) == 1 {
	//	return false
	//}
	//if len(s) == 0 {
	//	return true // 空字符串是符合的
	//}

	stack := make([]string, 0)

	for _, v := range []byte(s) {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			stack = append(stack, string(v))
		} else {
			pair := mapping[string(v)]
			if len(stack) > 0 && stack[len(stack)-1] == pair {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
