package main

import (
	"fmt"
	"strings"
)

/*
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。
如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指仅由字母组成、不包含任何空格的 最大子字符串。
示例:
输入: "Hello World"
输出: 5
*/

func main() {
	s := "Hello worddd"
	res := lengthOfLastWord(s)
	fmt.Println(res)
}

func lengthOfLastWord(s string) int {
	length := 0

	s = strings.Trim(s," ")
	for _, char := range s {
		if char != 32 {
			length++
		} else {
			length = 0
		}
	}

	return length
}
