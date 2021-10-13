package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSubarrayBoundedMax2([]int{2, 1, 4, 3}, 2, 3))
}

/*
给定一个元素都是正整数的数组A ，正整数 L 以及 R (L <= R)。
求连续、非空且其中最大元素满足大于等于L 小于等于R的子数组个数。
例如 :
输入:
A = [2, 1, 4, 3]
L = 2
R = 3
输出: 3
解释: 满足条件的子数组: [2], [2, 1], [3].
注意:
L, R  和 A[i] 都是整数，范围在 [0, 10^9]。
数组 A 的长度范围在[1, 50000]。
*/
// 注意，这里求的是连续且非空,要保证连续，用dfs是不好搞的
func numSubarrayBoundedMax(A []int, L int, R int) int {
	var ans int
	dfs(A, []int{}, 0, &ans, L, R)
	return ans
}
func dfs(nums, path []int, start int, ans *int, L int, R int) {
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		if check(path, L, R) {
			fmt.Println(path)
			*ans = *ans + 1
			dfs(nums, path, i+1, ans, L, R)
		}
		path = path[:len(path)-1]
	}
}
func check(path []int, L int, R int) bool {
	if len(path) <= 0 {
		return false
	}
	max := math.MinInt64
	for _, n := range path {
		if n > max {
			max = n
		}
	}
	return L <= max && R >= max
}

func numSubarrayBoundedMax2(A []int, L int, R int) int {
	return last(A, R) - last(A, L-1)
}

func last(nums []int, n int) int {
	ans, pre := 0, 0

	for _, i := range nums {
		if i > n {
			pre = 0
		} else {
			pre++
		}
		ans += pre
	}
	return ans
}
