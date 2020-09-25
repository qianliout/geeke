package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// ip := "2001:0db8:85a3:0:0:8A2E:0370:7334"
	ip := "2001:0db8:85a3:0000:0:8A2E:0370:733a"
	res := validIPAddress(ip)
	fmt.Println(res)
}

/*
编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址。
    如果是有效的 IPv4 地址，返回 "IPv4" ；
    如果是有效的 IPv6 地址，返回 "IPv6" ；
    如果不是上述类型的 IP 地址，返回 "Neither" 。
IPv4 地址由十进制数和点来表示，每个地址包含 4 个十进制数，其范围为 0 - 255， 用(".")分割。比如，172.16.254.1；
同时，IPv4 地址内的数不会以 0 开头。比如，地址 172.16.254.01 是不合法的。
IPv6 地址由 8 组 16 进制的数字来表示，每组表示 16 比特。这些组数字通过 (":")分割。比如,  2001:0db8:85a3:0000:0000:8a2e:0370:7334 是一个有效的地址。而且，我们可以加入一些以 0 开头的数字，字母可以使用大写，也可以是小写。所以， 2001:db8:85a3:0:0:8A2E:0370:7334 也是一个有效的 IPv6 address地址 (即，忽略 0 开头，忽略大小写)。
然而，我们不能因为某个组的值为 0，而使用一个空的组，以至于出现 (::) 的情况。 比如， 2001:0db8:85a3::8A2E:0370:7334 是无效的 IPv6 地址。
同时，在 IPv6 地址中，多余的 0 也是不被允许的。比如， 02001:0db8:85a3:0000:0000:8a2e:0370:7334 是无效的。
示例 1：
输入：IP = "172.16.254.1"
输出："IPv4"
解释：有效的 IPv4 地址，返回 "IPv4"
示例 2：
输入：IP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
输出："IPv6"
解释：有效的 IPv6 地址，返回 "IPv6"
示例 3：
输入：IP = "256.256.256.256"
输出："Neither"
解释：既不是 IPv4 地址，又不是 IPv6 地址
示例 4：
输入：IP = "2001:0db8:85a3:0:0:8A2E:0370:7334:"
输出："Neither"
示例 5：
输入：IP = "1e1.4.5.6"
输出："Neither"
提示：
    IP 仅由英文字母，数字，字符 '.' 和 ':' 组成。
*/

func validIPAddress(IP string) string {
	if strings.Contains(IP, ".") {
		if validIPV4(IP) {
			return "IPv4"
		} else {
			return "Neither"
		}
	} else if strings.Contains(IP, ":") {
		if validIPV6(IP) {
			return "IPv6"
		} else {
			return "Neither"
		}
	}
	return "Neither"
}

func validIPV4(IP string) bool {
	ss := strings.Split(IP, ".")
	if len(ss) != 4 {
		return false
	}
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			return false
		}
		if len(strconv.Itoa(n)) != len(ss[i]) {
			return false
		}
		if n >= 255 || n <= 0 {
			return false
		}
	}
	return true
}
func validIPV6(IP string) bool {
	s := strings.ToLower(IP)
	if strings.Contains(s, "::") {
		return false
	}
	ss := strings.Split(s, ":")
	if len(ss) != 8 {
		return false
	}

	for i := 0; i < len(ss); i++ {
		if !ishex(ss[i]) {
			return false
		}
		if len(ss[i]) == 0 || len(ss[i]) > 4 {
			return false
		}
	}
	return true
}

func ishex(s string) bool {
	for _, i := range []byte(s) {
		if _, is := fromHexChar(i); !is {
			return false
		}
	}
	return true

}

func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}
