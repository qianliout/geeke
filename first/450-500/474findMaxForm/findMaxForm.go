package main

import (
	"fmt"

	. "outback/leetcode/common"
)

func main() {
	array := []string{"10", "0001", "111001", "1", "0"}
	res := findMaxForm(array, 5, 3)
	fmt.Println("res is ", res)
}

/*
在计算机界中，我们总是追求用有限的资源获取最大的收益。
现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。
你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。
注意:
给定 0 和 1 的数量都不会超过 100。
给定字符串数组的长度不会超过 600。
示例 1:
输入: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
输出: 4
解释: 总共 4 个字符串可以通过 5 个 0 和 3 个 1 拼出，即 "10","0001","1","0" 。
示例 2:
输入: Array = {"10", "0", "1"}, m = 1, n = 1
输出: 2
解释: 你可以拼出 "10"，但之后就没有剩余数字了。更好的选择是拼出 "0" 和 "1" 。
*/

// fixme 最后一个测试用例,超过时间限制,是因为什么呢
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	word := make(map[string][]int)
	for _, w := range strs {
		one, zero := 0, 0
		for _, ch := range []byte(w) {
			if ch == '1' {
				one++
			} else if ch == '0' {
				zero++
			}
		}
		word[w] = []int{one, zero}
	}

	// dp[i][j] 表示使用i个0,和j个1,最后能拼出多少
	// 初始化
	dp := make(map[int]map[int]int)
	for i := 0; i <= m; i++ {
		dp[i] = make(map[int]int)
	}

	// 因为初值就是0,可以不赋值
	for _, w := range strs {
		one := word[w][0]
		zero := word[w][1]
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {

				if i-zero >= 0 && j-one >= 0 {
					dp[i][j] = Max(dp[i][j], dp[i-zero][j-one]+1)
				}
			}
		}
	}
	return dp[m][n]
}
