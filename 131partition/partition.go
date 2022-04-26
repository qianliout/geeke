package main

import (
	"fmt"
)

func main() {
	par := partition("aabccdcc")
	fmt.Println("par ", par)
}

/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
回文串 是正着读和反着读都一样的字符串。
*/
func partition(s string) [][]string {
	path, res, mem := make([]string, 0), make([][]string, 0), make(map[string]bool)
	backtrace([]byte(s), 0, len(s)-1, path, &res, mem)
	return res
}

func backtrace(ss []byte, start, end int, path []string, res *[][]string, mem map[string]bool) {
	if start > end {
		*res = append(*res, append([]string{}, path...))
		return
	}
	for i := start; i <= end; i++ {
		if !isPalindrome(ss[start:i+1], mem) {
			continue
		}
		path = append(path, string(ss[start:i+1]))
		backtrace(ss, i+1, end, path, res, mem)
		path = path[:len(path)-1]
	}
}

func isPalindrome(ss []byte, mem map[string]bool) bool {
	fmt.Println("len mam", len(mem))
	if b, ok := mem[string(ss)]; ok {
		return b
	}

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
			mem[string(ss)] = false
			return false
		}
		left++
		right--
	}
	mem[string(ss)] = true

	return true
}

func letter(ch byte) bool {
	if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
