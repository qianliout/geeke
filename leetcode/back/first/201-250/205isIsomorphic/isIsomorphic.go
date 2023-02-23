package main

import (
	"fmt"
	"strings"
)

func main() {
	isismo2("ab", "aa")
}

/*
给定两个字符串 s 和 t，判断它们是否是同构的。
如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
示例 1:
输入: s = "egg", t = "add"
输出: true
示例 2:
输入: s = "foo", t = "bar"
输出: false
示例 3:
输入: s = "paper", t = "title"
输出: true
说明:
你可以假设 s 和 t 具有相同的长度。
*/
// fixme 没有理解这种做法
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1 := make(map[byte]int)
	hash2 := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash1[s[i]] = i
		hash2[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		if hash2[t[i]] != hash1[s[i]] {
			return false
		}
	}
	return true
}

// 方法二,用map用两路比较
func isIsmo(s, t string) bool {
	return compareHelper(s, t) && compareHelper(t, s)
}

func compareHelper(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 从s到t
	map1 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		s1 := s[i]
		t1 := t[i]
		v, ok := map1[s1]
		if !ok {
			map1[s1] = t1
		} else {
			if v != t1 {
				return false
			}
		}
	}
	return true
}

func isismo2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return comHelper(s) == comHelper(t)
}

func comHelper(s string) string {
	count := 1
	exitMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := exitMap[s[i]]; !ok {
			exitMap[s[i]] = count
			count++
		}
	}
	nums := make([]int, 0)
	for i := 0; i < len(s); i++ {
		nums = append(nums, exitMap[s[i]])
	}
	res := strings.Replace(strings.Trim(fmt.Sprint(nums), "[]"), " ", "", -1)
	fmt.Println("res is ", res)
	return res
}
