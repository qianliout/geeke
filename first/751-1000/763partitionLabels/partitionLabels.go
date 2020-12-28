package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
示例：
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
*/
func partitionLabels(S string) []int {
	count := make(map[byte]int)
	ss := []byte(S)
	for _, s := range []byte(S) {
		count[s]++
	}
	res := make([]int, 0)
	widow := make([]byte, 0)
	i := 0
	for i < len(ss) {
		for !check(widow, count) {
			widow = append(widow, ss[i])
			i++
		}
		res = append(res, len(widow))
		widow = make([]byte, 0)
	}
	return res
}

func check(pres []byte, all map[byte]int) bool {
	if len(pres) <= 0 {
		return false
	}
	pre := make(map[byte]int)
	for _, k := range pres {
		pre[k]++
	}
	for k, v := range pre {
		if v < all[k] {
			return false
		}
	}
	return true
}
