package main

import (
	"fmt"
)

func main() {
	res := reverseVowels("leetcode")
	fmt.Println("res is ", res)
}

/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
示例 1:
输入: "hello"
输出: "holle"
示例 2:
输入: "leetcode"
输出: "leotcede"
*/

func reverseVowels(s string) string {
	if len(s) <= 1 {
		return s
	}
	ss := []byte(s)
	left, right := 0, len(ss)-1
	for left < right {
		i := left
		for i < right {
			if isVowels(ss[i]) {
				break
			}
			i++
		}
		j := right

		for j > left {
			if isVowels(ss[j]) {
				break
			}
			j--
		}
		if i < j {
			ss[i], ss[j] = ss[j], ss[i]
		}
		left = i + 1
		right = j - 1
	}
	return string(ss)
}

func isVowels(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'u' || c == 'o' || c == 'A' || c == 'E' || c == 'I' || c == 'U' || c == 'O' {
		return true
	}
	return false
}
