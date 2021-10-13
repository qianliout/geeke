package main

import (
	"fmt"
)

func main() {
	//s:="A man, a plan, a canal: Panama"
	s := "0P"
	res := isPalindrome(s)

	fmt.Println("res is ", res)
}

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:
输入: "race a car"
输出: false
*/
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := []byte(s)
	left, right := 0, len(ss)-1
	for left < right {
		if !isItALetterOrNumber(ss[left]) {
			left++
			continue
		}
		if !isItALetterOrNumber(ss[right]) {
			right--
			continue
		}
		if ToLow(ss[left]) != ToLow(ss[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isItALetterOrNumber(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || s >= '0' && s <= '9' {
		return true
	}
	return false
}

func ToLow(s byte) byte {
	if s >= 'A' && s <= 'Z' {
		s = s + 32
	}
	return s
}
