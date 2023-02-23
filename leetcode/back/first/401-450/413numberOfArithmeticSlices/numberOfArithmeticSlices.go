package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 12, 14, 16, 18}
	res := numberOfArithmeticSlices(nums)
	fmt.Println("res is ", res)

}

/*
如果一个数列至少有三个元素，并且任意两个相邻元素之差相同，则称该数列为等差数列。
例如，以下数列为等差数列:
1, 3, 5, 7, 9
7, 7, 7, 7
3, -1, -5, -9
以下数列不是等差数列。
1, 1, 2, 5, 7
数组 A 包含 N 个数，且索引从0开始。数组 A 的一个子数组划分为数组 (P, Q)，P 与 Q 是整数且满足 0<=P<Q<N 。
如果满足以下条件，则称子数组(P, Q)为等差数组：
元素 A[P], A[p + 1], ..., A[Q - 1], A[Q] 是等差的。并且 P + 1 < Q 。
函数要返回数组 A 中所有为等差数组的子数组个数。
示例:
A = [1, 2, 3, 4]
返回: 3, A 中有三个子等差数组: [1, 2, 3], [2, 3, 4] 以及自身 [1, 2, 3, 4]。
*/
// 方法一,三次遍历
func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}

	length := len(A)
	dif1 := make([]int, length-1)
	for i := 1; i < length; i++ {
		diff := A[i] - A[i-1]
		dif1[i-1] = diff

	}
	dif2 := make([]int, 0)

	r := 1
	for i := 1; i < len(dif1); i++ {
		if dif1[i] != dif1[i-1] {
			dif2 = append(dif2, r)
			r = 1
		} else {
			r++
		}
	}
	dif2 = append(dif2, r)

	res := 0
	for _, i := range dif2 {
		if i >= 2 {
			res += (i) * (i - 1) / 2
		}
	}
	return res
}

func numberOfArithmeticSlices2(A []int) int {
	// dp[i]表示以下标i结尾的等差数列的个数
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 0
	sum := 0
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1 // 这里认真理解
			sum += dp[i]
		} else {
			dp[i] = 0
		}
	}
	return sum
}
