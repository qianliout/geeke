package main

import (
	"fmt"
)

func main() {
	res := isValid("()[]{}")
	fmt.Println(res)
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
示例 1：
输入：s = "()"
输出：true
*/
func isValid(s string) bool {
	match := map[byte]byte{'}': '{', ')': '(', ']': '['}
	
	stark := make([]byte, 0)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ss[i] == '(' || ss[i] == '{' || ss[i] == '[' {
			stark = append(stark, ss[i])
		} else {
			if len(stark) == 0 {
				return false
			}
			if stark[len(stark)-1] != match[ss[i]] {
				return false
			}
			stark = stark[:len(stark)-1]
		}
	}
	return len(stark) == 0
}
