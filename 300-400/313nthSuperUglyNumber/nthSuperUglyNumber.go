package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/common/utils"
)

func main() {
	n := 12
	primes := []int{2, 7, 13, 19}
	number := nthSuperUglyNumber(n, primes)
	fmt.Println("numbers ", number)

}

func nthSuperUglyNumber(n int, primes []int) int {
	// dp[i]表示第i个丑数的值
	dp := make([]int, n)
	// idx[i]表示：primes[i]*dp[idx[i]] 可以生成下一个候选的丑书
	idx := make([]int, len(primes))
	dp[0] = 1
	for i := 1; i < n; i++ {
		ugly := math.MaxInt
		for k := 0; k < len(primes); k++ {
			ugly = Min(ugly, primes[k]*dp[idx[k]])
		}
		dp[i] = ugly
		for k := range primes {
			if dp[i] == primes[k]*dp[idx[k]] {
				idx[k]++
			}
		}
	}
	return dp[n-1]
}
