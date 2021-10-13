package main

import (
	"fmt"
)

func main() {
	res := countNumbersWithUniqueDigits(3)
	fmt.Println("res is ", res)
}

/*
给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10n 。
示例:
输入: 2
输出: 91
解释: 答案应为除去 11,22,33,44,55,66,77,88,99 外，在 [0,100) 区间内的所有数字。
*/
func countNumbersWithUniqueDigits(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 1, 10
	if n <= 1 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		num := 9
		a := 9
		for j := i - 1; j > 0; j-- {
			num = num * a
			a--
		}
		dp[i] = dp[i-1] + num
	}
	return dp[n]
}
