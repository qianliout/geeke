package main

import (
	"fmt"
)

func main() {
	fmt.Println(reorganizeString("kkkkzrkatkwpkkkktrq"))
	fmt.Println(reorganizeString("todrnphcamnomskfrhe"))
	fmt.Println(reorganizeString("aba"))
	fmt.Println(reorganizeString("abab"))
}

/*
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
若可行，输出任意可行的结果。若不可行，返回空字符串。
示例 1:
输入: S = "aab"
输出: "aba"
示例 2:
输入: S = "aaab"
输出: ""
注意:
S 只包含小写字母并且长度在[1, 500]区间内。
*/
func reorganizeString(S string) string {
	ss := []byte(S)
	countMap := make(map[byte]int)
	for _, s := range ss {
		countMap[s]++
	}
	res := make([]byte, 0)
	dfs(countMap, []byte(""), len(ss), &res)
	return string(res)
}

// dfs 搞不定，因为找不到退出条件
func dfs(countMap map[byte]int, path []byte, l int, res *[]byte) {
	if len(path) == l {
		*res = append([]byte{}, path...)
		return
	}
	for k, v := range countMap {
		if v > 0 && (len(path) == 0 || (len(path) > 0 && (path)[len(path)-1] != k)) {
			path = append(path, k)
			countMap[k]--
			dfs(countMap, path, l, res)
			path = path[:len(path)-1]
			countMap[k]++
		}
	}
}
