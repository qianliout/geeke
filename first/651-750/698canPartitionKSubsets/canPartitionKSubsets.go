package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{815, 625, 3889, 4471, 60, 494, 944, 1118, 4623, 497, 771, 679, 1240, 202, 601, 883}, 3))
}

/*
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
示例 1：
输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
提示：
1 <= k <= len(nums) <= 16
0 < nums[i] < 10000
*/
//  方法正确，但是会超时
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) <= 0 {
		return false
	}
	sum := 0
	max := math.MinInt64
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if sum%k != 0 || sum/k < max {
		return false
	}
	target := sum / k
	used := make(map[int]bool)
	return backtracking(nums, target, k, 0, used)
}

func backtracking(nums []int, target, k int, path int, used map[int]bool) bool {
	if k <= 0 {
		return true
	}
	if path == target {
		return backtracking(nums, target, k-1, 0, used)
	}

	for i, n := range nums {
		if !used[i] && n+path <= target {
			used[i] = true
			if backtracking(nums, target, k, path+n, used) {
				return true
			}
			used[i] = false
		}
	}
	return false
}
