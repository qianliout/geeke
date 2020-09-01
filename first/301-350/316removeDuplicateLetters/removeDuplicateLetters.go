package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "daaabbbccca"
	res := removeDuplicateLetters2(s)
	fmt.Println("res is ", res)
}

/*
给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
示例 1:
输入: "bcabc"
输出: "abc"
示例 2:
输入: "cbacdcbc"
输出: "acdb"
链接：https://leetcode-cn.com/problems/remove-duplicate-letters
*/
// 使用stark的方式
func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}
	exit := make(map[byte]int)
	for _, i := range []byte(s) {
		exit[i] += 1
	}
	exit2 := make(map[byte]bool)

	stark := make([]byte, 0)

	for _, i := range []byte(s) {
		if exit2[i] {
			exit[i] -= 1
			continue // 已以在里面的了
		}
		if len(stark) == 0 || (len(stark) > 0 && stark[len(stark)-1] < i) {
			stark = append(stark, i)
			exit[i] -= 1
			exit2[i] = true
		} else {
			for len(stark) > 0 && stark[len(stark)-1] > i {
				peek := stark[len(stark)-1]
				if exit[peek] <= 0 {
					break
				}
				exit2[peek] = false
				stark = stark[:len(stark)-1]
			}
			stark = append(stark, i)
			exit[i] -= 1
			exit2[i] = true
		}
	}
	res := ""
	for _, i := range stark {
		res = res + string(i)
	}
	return res
}

// 使用递归，贪心算法
func removeDuplicateLetters2(s string) string {
	if len(s) == 0 {
		return ""
	}
	exit := make(map[byte]int)
	for _, i := range []byte(s) {
		exit[i] += 1
	}
	pos := 0
	for i, v := range []byte(s) {
		if v < []byte(s)[pos] {
			pos = i
		}
		exit[v]--
		if exit[v] == 0 {
			break
		}
	}
	pre := string([]byte(s)[pos])
	new2 := strings.Replace(string([]byte(s)[pos:]), pre, "", -1)
	return pre + removeDuplicateLetters2(new2)
}
