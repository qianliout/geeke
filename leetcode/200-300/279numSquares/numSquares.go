package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	fmt.Println(numSquares(-1))
	// fmt.Println(numSquares(13))
	// fmt.Println(numSquares(15))
	// fmt.Println(numSquares(16))
}

// 贪心算法还不得行，还只能用dp
func numSquares(n int) int {
	dp := make(map[int]int)

	for i := 1; i <= n; i++ {
		// 初值
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func find(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return int(math.Log2(float64(n)))
}
