package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findSubstringInWraproundString("zab"))
}

/*
把字符串 s 看作是“abcdefghijklmnopqrstuvwxyz”的无限环绕字符串，所以 s 看起来是这样的："...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd....".
现在我们有了另一个字符串 p 。你需要的是找出 s 中有多少个唯一的 p 的非空子串，尤其是当你的输入是字符串 p ，你需要输出字符串 s 中 p 的不同的非空子串的数目。
注意: p 仅由小写的英文字母组成，p 的大小可能超过 10000。
示例 1:
输入: "a"
输出: 1
解释: 字符串 S 中只有一个"a"子字符。
示例 2:
输入: "cac"
输出: 2
解释: 字符串 S 中的字符串“cac”只有两个子串“a”、“c”。.
示例 3:
输入: "zab"
输出: 6
解释: 在字符串 S 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。.
*/
// dp表示以当前字符结尾的最长递增子串的长度，map数组里存的是当前状态以a-z字母结尾的最长子串长度。
// 每访问一个字符，则首先更新dp值，连续的则dp+1， 否则dp等于1。然后将dp值与map里对应的值比较，
// 大于的话说明有新的以当前字母结尾的子串，更新sum的值。
// 理解dp的含义就知道，后面求值时为什么可以直接用加法了，比如dp[e]=5 ,也就是说abcde,出在在s中，但是此时，dp[e]=5,表示的是，abcde,bcde,cde,de,e 这五个数，
// 而不用计算前面，因为前面的有dp[d],dp[c],dp[b],de[a]来计算

func findSubstringInWraproundString(p string) int {
	if len(p) <= 1 {
		return len(p)
	}
	ss := []byte(p)
	dp := make([]int, 26) // 26个字母
	cnt := 1
	for i := 0; i < len(p); i++ {
		idx := ss[i] - 'a'

		if i-1 >= 0 {
			if check(ss[i-1], ss[i]) {
				cnt++
			} else {
				cnt = 1
			}
		}
		dp[idx] = Max(dp[idx], cnt)
	}
	ans := 0
	for _, n := range dp {
		ans += n
	}
	return ans
}
func check(a, b byte) bool {
	if a == 'z' && b == 'a' {
		return true
	}
	return b-a == 1
}
func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
