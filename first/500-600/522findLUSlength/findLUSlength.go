package main

import (
	"fmt"
)

func main() {
	strc := []string{"aaa", "aaa", "aa"}
	res := findLUSlength(strc)
	fmt.Println("res is ", res)
}

/*
给定字符串列表，你需要从它们中找出最长的特殊序列。最长特殊序列定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。
子序列可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。
输入将是一个字符串列表，输出是最长特殊序列的长度。如果最长特殊序列不存在，返回 -1 。
示例：
输入: "aba", "cdc", "eae"
输出: 3
提示：
    所有给定的字符串长度不会超过 10 。
    给定字符串列表的长度将在 [2, 50 ] 之间。
*/

// 1.把所有重复的字符（用list2记录重复字符串，后面需要使用）串删除。
// 2.把第一步删除后的列表中所有属于list2的子字符串删除。（这一步较难）
// 3.剩下都是属于特殊序列，只需要找出最长的即可。
func findLUSlength(strs []string) int {
	ans := -1
	hash := make(map[string]int)  // 记录重复的
	hash2 := make(map[string]int) // 记录子串

	for _, s := range strs {
		hash[s]++
	}
	for i := 0; i < len(strs); i++ {
		if hash[strs[i]] > 1 {
			continue
		}
		for k, v := range hash {
			if v <= 1 {
				continue
			}
			if lus(strs[i], k) {
				hash2[strs[i]]++
			}
		}
	}
	for _, k := range strs {
		if hash[k] <= 1 && hash2[k] == 0 && len(k) > ans {
			ans = len(k)
		}
	}

	return ans
}

// 判断s1不是不s2的子序列
func lus(s1, s2 string) bool {
	if s1 == "" || s1 == s2 {
		return true
	}
	if len(s1) > len(s2) {
		return false
	}
	ll, rl := 0, 0
	for ll < len(s1) && rl < len(s2) {
		for rl < len(s2) {
			if s1[ll] == s2[rl] {
				ll++
				rl++
				break
			}
			rl++
		}
	}
	return ll == len(s1)
}

func findLus(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
