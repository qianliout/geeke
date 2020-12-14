package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("abc"))
}

/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
示例 1：
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
提示：
输入的字符串长度不会超过 1000 。
*/
// 当只有一个字符时，比如 a 自然是一个回文串。
// 当有两个字符时，如果是相等的，比如 aa，也是一个回文串。
// 当有三个及以上字符时，比如 ababa 这个字符记作串 1，把两边的 a 去掉，也就是 bab 记作串 2，可以看出只要串2是一个回文串，那么左右各多了一个 a 的串 1 必定也是回文串。所以当 s[i]==s[j] 时，自然要看 dp[i+1][j-1] 是不是一个回文串。
// 很好的一道题
func countSubstrings(s string) int {
	ss := []byte(s)
	// dp[i][j]表示以i为起点的个数
	dp := make(map[int]map[int]bool)
	for i := 0; i < len(ss); i++ {
		dp[i] = make(map[int]bool)
		dp[i][i] = true // 这一句也可以没有
	}
	ans := 0
	for j := 0; j < len(ss); j++ {
		for i := 0; i <= j; i++ {
			// 这里j-i<2怎么理解呢，j-i<2表示只有两个字符
			if ss[i] == ss[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				ans++
			}
		}
	}
	return ans
}
