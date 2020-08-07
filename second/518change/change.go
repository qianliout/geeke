package main

import (
	"fmt"
)

func main() {
	coins := []int{1, 2, 5}
	res := change(5, coins)
	fmt.Println(res)
}

func change(amount int, coins []int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	dp := make(map[int]int)
	dp[0] = 1
	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}

	return dp[amount]
}
