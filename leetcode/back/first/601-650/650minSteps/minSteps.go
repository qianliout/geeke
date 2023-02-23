package main

import (
	"math"
)

func main() {

}

/*
最初在一个记事本上只有一个字符 'A'。你每次可以对这个记事本进行两种操作：
Copy All (复制全部) : 你可以复制这个记事本中的所有字符(部分的复制是不允许的)。
Paste (粘贴) : 你可以粘贴你上一次复制的字符。
给定一个数字 n 。你需要使用最少的操作次数，在记事本中打印出恰好 n 个 'A'。输出能够打印出 n 个 'A' 的最少操作次数。
示例 1:
输入: 3
输出: 3
解释:
最初, 我们只有一个字符 'A'。
第 1 步, 我们使用 Copy All 操作。
第 2 步, 我们使用 Paste 操作来获得 'AA'。
第 3 步, 我们使用 Paste 操作来获得 'AAA'。
*/
// https://leetcode-cn.com/problems/2-keys-keyboard/solution/shu-wo-zhi-yan-bu-shao-ti-jie-gen-ben-bu-dong-dpi-/
// 素数只能是一个一个的粘贴， 合数可以由最大的质因数得来，本质就是素数的最大质因数
func minSteps(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= n; i++ {
		dp[i] = i
		// 因为是找最大的质因数，可以从后往前找
		for j := i / 2; i >= int(math.Sqrt(float64(i))); j-- {
			if i%j == 0 {
				dp[i] = dp[i/j] + dp[j] // 此时j,
				break
			}
		}
	}
	return dp[n]
}
