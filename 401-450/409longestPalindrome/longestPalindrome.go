package main

func main() {

}

/*
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
注意:
假设字符串的长度不会超过 1010。
示例 1:
输入:
"abccccdd"
输出:
7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
*/

// 这里是构造成的最大回方串
func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[int32]int)
	for _, v := range s {
		m[v]++
	}
	ans := 0
	maxOdd := 0

	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else {
			if v > maxOdd {
				maxOdd = v
			}
			ans += v - 1
		}
	}

	if maxOdd > 0 {
		ans += 1
	}
	return ans
}
