package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums := []int{4, 3, 2, 6}
	res := maxRotateFunction(nums)
	fmt.Println("res is ", res)
}

/*
给定一个长度为 n 的整数数组 A 。
假设 Bk 是数组 A 顺时针旋转 k 个位置后的数组，我们定义 A 的“旋转函数” F 为：
F(k) = 0 * Bk[0] + 1 * Bk[1] + ... + (n-1) * Bk[n-1]。
计算F(0), F(1), ..., F(n-1)中的最大值。
注意:
可以认为 n 的值小于 105。
示例:
A = [4, 3, 2, 6]
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。
*/
func maxRotateFunction(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	dp := make(map[int]int)
	length := len(A)
	dp[0] = F(A)

	max := dp[0]
	for i := 1; i < length; i++ {
		A = append([]int{A[length-1]}, A[:length-1]...)
		n := addAll(A[1:length]) + (1-length)*A[0]
		dp[i] = n + dp[i-1]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func maxRotateFunction2(A []int) int {
	if len(A) <= 1 {
		return 0
	}
	dp := make(map[int]int)
	length := len(A)
	dp[0] = F(A)
	sum := addAll(A)

	max := dp[0]
	for i := 1; i < length; i++ {
		n := sum - A[length-i] + (1-length)*A[length-i]
		dp[i] = n + dp[i-1]
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func addAll(nums []int) int {
	res := 0
	for _, nu := range nums {
		res += nu
	}
	return res
}

func F(nums []int) int {
	res := 0
	for i, nu := range nums {
		res += i * nu
	}
	return res
}
