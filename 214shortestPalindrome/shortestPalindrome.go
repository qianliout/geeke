package main

import (
	"fmt"
)

func main() {
	s := "ananab"
	palindrome := shortestPalindrome(s)
	fmt.Println(palindrome)
}

/*
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
*/
func shortestPalindrome(s string) string {
	ss1 := []byte(s)
	ss2 := reverse(ss1)
	notSame := findNotSame(ss2)
	return string(append(notSame, ss1...))
}

func reverse(ss []byte) []byte {
	ans := make([]byte, len(ss))
	for i := len(ss) - 1; i >= 0; i-- {
		ans[len(ss)-1-i] = ss[i]
	}
	return ans
}

func findNotSame(ss []byte) []byte {
	i := 0
	for i < len(ss) {

		if !palindrome(ss[i:]) {
			i++
		} else {
			break
		}
	}
	return ss[:i]
}

func palindrome(ss []byte) bool {
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}
	return true
}
