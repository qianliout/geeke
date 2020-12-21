package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkValidString("(())((())()()(*)(*()(())())())()()((()())((()))(*"))
	// fmt.Println(checkValidString("((((*)(*(((((*"))
}

/*
给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：
任何左括号 ( 必须有相应的右括号 )。
任何右括号 ) 必须有相应的左括号 ( 。
左括号 ( 必须在对应的右括号之前 )。
* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
一个空字符串也被视为有效字符串。
示例 1:
输入: "()"
输出: True
示例 2:
输入: "(*)"
输出: True
示例 3:
输入: "(*))"
输出: True
*/
func checkValidString(s string) bool {
	stark1 := make([]int, 0)
	stark2 := make([]int, 0)
	ss := []byte(s)
	for i, b := range ss {
		if b == '(' {
			stark1 = append(stark1, i)
		} else if b == '*' {
			stark2 = append(stark2, i)
		} else if b == ')' {
			if len(stark1) > 0 {
				stark1 = stark1[:len(stark1)-1]
			} else if len(stark1) == 0 && len(stark2) > 0 {
				stark2 = stark2[:len(stark2)-1]
			} else {
				return false
			}
		}
	}
	if len(stark1) > len(stark2) {
		return false
	}
	// 这一步判断是最容易忘的
	for len(stark1) > 0 {
		if stark1[len(stark1)-1] > stark2[len(stark2)-1] {
			return false
		}
		stark1 = stark1[:len(stark1)-1]
		stark2 = stark2[:len(stark2)-1]
	}

	return true
}
