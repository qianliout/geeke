package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{546, 669}
	//nums := []int{4, 8, 10, 240}
	nums := []int{1, 2, 4, 8}
	res := largestDivisibleSubset(nums)
	fmt.Println("res is ", res)
}

/*
给出一个由无重复的正整数组成的集合，找出其中最大的整除子集，子集中任意一对 (Si，Sj) 都要满足：Si % Sj = 0 或 Sj % Si = 0。
如果有多个目标子集，返回其中任何一个均可。
示例 1:
输入: [1,2,3]
输出: [1,2] (当然, [1,3] 也正确)
示例 2:
输入: [1,2,4,8]
输出: [1,2,4,8]
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	if len(nums) <= 1 {
		return nums
	}
	// 题目中说一定是正整数
	dp := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		dp[i] = []int{nums[i]}
	}
	res := dp[0]

	for i := 1; i < len(nums); i++ {
		pre := dp[i]
		preI := i
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(dp[j]) >= len(pre) {
					pre = dp[j]
					preI = j
				}
			}
		}
		if preI != i {
			dp[i] = append(pre, nums[i])
		}
		if len(res) < len(dp[i]) {
			res = dp[i]
		}
	}
	return res
}
