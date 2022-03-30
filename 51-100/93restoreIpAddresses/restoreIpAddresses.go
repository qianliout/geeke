package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ips := restoreIpAddresses("25525511135")
	fmt.Println("ips is ", ips)
}

/*
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
*/

func restoreIpAddresses(s string) []string {
	ss := []byte(s)
	path := make([]string, 0)
	res := make([]string, 0)
	dfs(ss, 0, path, &res)
	return res
}

func dfs(ss []byte, start int, path []string, res *[]string) {
	if len(path) == 4 {
		ip := strings.Join(path, ".")
		if len(ip)-3 == len(ss) {
			*res = append(*res, ip)
		}
		return
	}
	if start > len(ss) {
		return
	}
	for i := start + 1; i <= len(ss); i++ {
		if check(string(ss[start:i])) {
			path = append(path, string(ss[start:i]))
			dfs(ss, i, path, res)
			path = path[:len(path)-1]
		}
	}
}

func check(ip string) bool {
	parseInt, err := strconv.ParseInt(ip, 10, 64)
	if err != nil {
		return false
	}
	if parseInt < 0 || parseInt > 255 {
		return false
	}
	if strconv.Itoa(int(parseInt)) != ip {
		return false // 说明有前导0或多个0
	}
	return true
}
