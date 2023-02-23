package main

import (
	"fmt"
	"strings"
)

func main() {
	res := restoreIpAddresses("0000")
	fmt.Println("res si ", res)
}

/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
有效的 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效的 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效的 IP 地址。
示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：
输入：s = "0000"
输出：["0.0.0.0"]
示例 3：
输入：s = "1111"
输出：["1.1.1.1"]
示例 4：
输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
示例 5：
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
提示：
    0 <= s.length <= 3000
    s 仅由数字组成
*/
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)
	dfs([]byte(s), 0, 4, path, &res)
	return res
}

// residue表示还有多少个段待分配
func dfs(ss []byte, start, residue int, path []string, res *[]string) {
	if start == len(ss) {
		if residue == 0 {
			*res = append(*res, strings.Join(path, "."))
		}
		return
	}
	for i := start; i < start+3; i++ {
		if i >= len(ss) {
			break
		}
		if residue*3 < len(ss)-i {
			continue
		}
		if ipIsTrue(ss, start, i) {
			path = append(path, string(ss[start:i+1]))
			dfs(ss, i+1, residue-1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func ipIsTrue(ss []byte, left, right int) bool {
	length := right - left + 1
	if length > 1 && ss[left] == '0' {
		return false
	}

	res := 0
	for left <= right {
		res = res*10 + int(ss[left]-'0')
		left++
	}
	if res < 0 || res > 255 {
		return false
	}
	return true
}
