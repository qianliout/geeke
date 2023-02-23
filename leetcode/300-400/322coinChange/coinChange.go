package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	coins := []int{2}
	amount := 3

	res := coinChange(coins, amount)

	fmt.Println("res is ", res)

}

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)

	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			} else if i-coins[j] == 0 {
				dp[i] = 1
			} else {
				dp[i] = Min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
