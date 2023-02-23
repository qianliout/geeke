package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "   fly me   to   the moon  "
	s := "moon"
	lastWord := lengthOfLastWord(s)
	fmt.Println("last word ", lastWord)
}

/*
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中最后一个单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
示例 1：
输入：s = "Hello World"
输出：5
示例 2：
输入：s = "   fly me   to   the moon  "
输出：4
示例 3：
输入：s = "luffy is still joyboy"
输出：6
*/
func lengthOfLastWord1(s string) int {
	right, first := len(s)-1, false
	ss := []byte(s)
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == ' ' {
			if !first {
				continue
			} else {
				return right - i
			}
		} else {
			if !first {
				right = i
			}
			first = true
		}
	}
	return right + 1
}

func lengthOfLastWord(s string) int {
	ss := []byte(strings.Trim(s, " "))
	right := len(ss) - 1
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == ' ' {
			return right - i
		}
	}
	return len(ss)
}
