package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "A man, a plan, a canal: Panama"
	s := ""
	palindrome := isPalindrome(s)
	fmt.Println("palindrome", palindrome)
}

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
*/
func isPalindrome(s string) bool {
	ss := []byte(strings.ToLower(s))
	left, right := 0, len(ss)-1
	for left <= right {
		if !letter(ss[left]) {
			left++
			continue
		}
		if !letter(ss[right]) {
			right--
			continue
		}
		if ss[left] != ss[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func letter(ch byte) bool {
	if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
