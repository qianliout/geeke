package main

import (
	"outback/leetcode/back/common"
)

func main() {

}

/*
给定一个整数数组 nums ，你可以对它进行一些操作。
每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除每个等于 nums[i] - 1 或 nums[i] + 1 的元素。
开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。
示例 1:
输入: nums = [3, 4, 2]
输出: 6
解释:
删除 4 来获得 4 个点数，因此 3 也被删除。
之后，删除 2 来获得 2 个点数。总共获得 6 个点数。
示例 2:
输入: nums = [2, 2, 3, 3, 3, 4]
输出: 9
解释:
删除 3 来获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。
注意:
nums的长度最大为20000。
每个整数nums[i]的大小都在[1, 10000]范围内。
*/
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	for _, i := range nums {
		if i > max {
			max = i
		}
	}
	all := make([]int, max+1)
	for _, i := range nums {
		all[i]++
	}
	dp := make(map[int]int)
	dp[1] = all[1]
	dp[2] = common.Max(dp[1], all[2]*2)
	// dp转移方程
	// 如果你不删除当前位置的数字，那么你得到就是前一个数字的位置的最优结果。
	// 如果你觉得当前的位置数字i需要被删，那么你就会得到i - 2位置的那个最优结果加上当前位置的数字乘以个数。
	for i := 2; i <= max; i++ {
		dp[i] = common.Max(dp[i-1], dp[i-2]+all[i]*i)
	}
	return dp[max]
}
