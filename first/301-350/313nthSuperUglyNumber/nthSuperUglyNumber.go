package main

import (
	"fmt"
	"math"
)

func main() {
	n := 500
	primes := []int{5, 7, 13, 17, 23, 29, 31, 43, 53, 59, 61, 71, 73, 79, 83, 97, 109, 131, 137, 163, 167, 181, 191, 193, 197, 199, 227, 233, 251, 263}
	res := nthSuperUglyNumber(n, primes)
	fmt.Println("res is ", res)
}

/*
编写一段程序来查找第 n 个超级丑数。
超级丑数是指其所有质因数都是长度为 k 的质数列表 primes 中的正整数。
示例:
输入: n = 12, primes = [2,7,13,19]
输出: 32
解释: 给定长度为 4 的质数列表 primes = [2,7,13,19]，前 12 个超级丑数序列为：[1,2,4,7,8,13,14,16,19,26,28,32] 。
说明:
    1 是任何给定 primes 的超级丑数。
     给定 primes 中的数字以升序排列。
    0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000 。
    第 n 个超级丑数确保在 32 位有符整数范围内。
链接：https://leetcode-cn.com/problems/super-ugly-number
*/

// 自已在本地不会超出时间限制，但是提交会超出限制
// 使用小顶堆
func nthSuperUglyNumber2(n int, primes []int) int {
	return 0
}

func nthSuperUglyNumber(n int, primes []int) int {
	idx := make([]int, len(primes))
	// 理解这个idx表示的是什么是这道题的关键，以题目的样例为例
	/*
			i等于1时，dp[i] = primes[0]*dp[0] 也就是2*1  此时idx[0] = 1 (原来是0，dp[i],选中的是primes[0],所以idx[0]要加一)
			i等于2时，dp[i] =  min(dp[0]*idx[0],dp[1]*idx[1]....) 显然应该是dp[0]*idx[0] = 2*2 (从这里也可以看出为什么idx[k]++)
		因为不加加的话就会再次出现一个2，导致重复计算，得不到正确的值
	*/
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		min := math.MaxInt32
		for k := 0; k < len(primes); k++ {
			tmp := primes[k] * dp[idx[k]]
			if tmp < min {
				min = tmp
			}
		}
		dp[i] = min

		for k := 0; k < len(primes); k++ {
			tmp := primes[k] * dp[idx[k]]
			if tmp == min {
				idx[k]++ // 下次再使用于时就要使用这个增加后的
			}
		}
	}
	//fmt.Println(dp)
	return dp[n-1]
}

func nthSuperUglyNumber3(n int, primes []int) int {
	idx := make([]int, len(primes))
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		min := math.MaxInt32
		for k := 0; k < len(primes); k++ {
			tmp := primes[k] * dp[idx[k]]
			if tmp < min {
				min = tmp
			}
		}
		dp[i] = min
		for k := 0; k < len(primes); k++ {
			tmp := primes[k] * dp[idx[k]]
			if tmp == min {
				idx[k]++
			}
		}
	}
	//fmt.Println(dp)
	return dp[n-1]
}
