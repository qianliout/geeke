package main

import (
	"fmt"
	"math"
)

func main() {
	res := rangeBitwiseAnd(5, 7)
	fmt.Println("res is ", res)
}

/*
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。
示例 1:
输入: [5,7]
输出: 4
示例 2:
输入: [0,1]
输出: 0
链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range
*/

func rangeBitwiseAnd(m int, n int) int {
	return find3(m, n)
}

// 结果正确，但是会超时
func iteration(m, n int) int {
	res := m

	for i := m + 1; i <= n; i++ {
		res = res & i
		// 这里一定要注意，当i=2147483646时（mathInt32）再加一就变成负数，又进入到循环了，这循环就不退出了
		if res == 0 || res == math.MaxInt32 {
			break
		}
	}
	return res
}

// https://leetcode-cn.com/problems/bitwise-and-of-numbers-range/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--41/
// 只要有一个为0，那么他和其他所有的数按位与都是0,所以这道题就变成了求公共的位前缀
func find2(m, n int) int {
	right := 0

	for m < n {
		m >>= 1
		n >>= 1
		right++
	}
	return m << right
}

func find3(m, n int) int {
	for m < n {
		n = n & (n - 1) // 这个方法是打掉最后的一个1
	}
	return n
}

// 还有一个解法3，还没有懂
