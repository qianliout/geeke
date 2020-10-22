package main

import (
	"fmt"
)

func main() {
	nums := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}
	res := canPartition2(nums)
	fmt.Println("res is", res)
}

/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
注意:
    每个数组中的元素不会超过 100
    数组的大小不会超过 200
示例 1:
输入: [1, 5, 11, 5]
输出: true
解释: 数组可以分割成 [1, 5, 5] 和 [11].
示例 2:
输入: [1, 2, 3, 5]
输出: false
解释: 数组不能分割成两个元素和相等的子集.
*/

// 一维dp,状态压缩
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	dp := make(map[int]bool)
	// base case
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum / 2; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}
	return dp[sum/2]
}

// 二维dp，可以得到结果，但是会超时，要进行状态压缩
func canPartition2(nums []int) bool {
	// 定义dp dp[i][j] 表示选择第i个物品，当前容量是j时是否可以刚好装满
	dp := make(map[int]map[int]bool)
	// 初始化并赋初值
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	// 初始化，初值
	for i := 0; i <= sum/2; i++ {
		if dp[0] == nil {
			dp[0] = make(map[int]bool)
		}
		dp[0][i] = false
	}
	for i := 0; i <= len(nums); i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]bool)
		}
		dp[i][0] = true
	}
	// 状态转移
	// dp[i-1][j]表示第i个物品不放进去，那个此时就前i-1时是否装满
	// dp[i-1][j-nums[i-1]] 表示第i个物品放进去，那么，就看j-nums[i-1]容量时是否放满
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum/2; j++ {
			if j-nums[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				// 背包容量不足，不能装入第 i 个物品
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][sum/2]
}
