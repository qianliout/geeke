package main

import (
	"fmt"
)

func main() {
	res := isScramble("great", "rgeae")
	fmt.Println("res is ", res)
}

/*
给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
下图是字符串 s1 = "great" 的一种可能的表示形式。
    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t
在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t
我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a
我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
示例 1:
输入: s1 = "great", s2 = "rgeat"
输出: true
示例 2:
输入: s1 = "abcde", s2 = "caebd"
输出: false
*/

// 注意,这题可以随便从中间分,而不仅仅从中间分,这里使用三维dp
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	ss1 := []byte(s1)
	ss2 := []byte(s2)

	n := len(s1)
	dp := make(map[int]map[int]map[int]bool)

	// 初值
	for i := 0; i <= n; i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]map[int]bool)
		}
		for j := 0; j <= n; j++ {
			if dp[i][j] == nil {
				dp[i][j] = make(map[int]bool)
			}
		}
	}

	// 状态方程
	// 这里的l表法起点开始的长度,假设起点是0,长度是1.,那就只有第0个元素,是个右开区间,所以从1开始到n
	for l := 1; l <= n; l++ {
		// S1开始的地方
		for i := 0; i+l <= n; i++ {
			// S2的开始
			for j := 0; j+l <= n; j++ {
				// 如果长度是1,那么是没法分隔的,就只比较原数据就可以了
				if l == 1 {
					dp[l][i][j] = ss1[i] == ss2[j]
				} else {
					// 如果长度不是1,那就要分隔了,
					for p := 1; p < l; p++ {
						dp[l][i][j] = (dp[p][i][j] && dp[l-p][i+p][j+p]) || (dp[p][i][j+l-p] && dp[l-p][i+p][j])
						if dp[l][i][j] == true {
							break
						}
					}
				}
			}
		}

	}
	return dp[n][0][0]
}
