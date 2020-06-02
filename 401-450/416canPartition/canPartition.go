package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 5}
	res := canPartition(nums)
	fmt.Println("res is ", res)
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
示例 2:
输入: [1, 2, 3, 5]
输出: false
解释: 数组不能分割成两个元素和相等的子集.
*/
// 方法一
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	dp := make(map[int]bool)
	//

	for _, num := range nums {
		for i := sum / 2; i >= 0; i-- {
			if i == num {
				dp[i] = true
			} else if i-num > 0 {
				dp[i] = dp[i] || dp[i-num]
			}
		}
	}
	//fmt.Println(dp)
	return dp[sum/2]
}
