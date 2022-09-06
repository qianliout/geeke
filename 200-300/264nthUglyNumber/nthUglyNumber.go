package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	number := nthUglyNumber(10)
	fmt.Println("res ", number)
}

/*
给你一个整数 n ，请你找出并返回第 n 个 丑数 。
丑数 就是只包含质因数 2、3 和/或 5 的正整数。
*/
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	// dp 表示第n个丑数的值
	dp, exit := make(map[int]int), make(map[int]bool)
	dp[0] = 1
	i, i2, i3, i5 := 1, 0, 0, 0
	for i <= n {
		tem := Min(dp[i2]*2, dp[i3]*3, dp[i5]*5)
		if !exit[tem] {
			dp[i] = tem
			exit[tem] = true
			i++
		}
		if tem == dp[i2]*2 {
			i2++
		} else if tem == dp[i3]*3 {
			i3++
		} else if tem == dp[i5]*5 {
			i5++
		}
	}
	return dp[n-1]
}
