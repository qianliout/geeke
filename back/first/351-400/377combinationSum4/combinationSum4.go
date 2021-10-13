package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := combinationSum42(nums, 4)
	fmt.Println("res is ", res)
}

/*
给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
示例:
nums = [1, 2, 3]
target = 4
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。
因此输出为 7。
进阶：
如果给定的数组中含有负数会怎么样？
问题会产生什么变化？
我们需要在题目中添加什么限制来允许负数的出现？
*/

func combinationSum4(nums []int, target int) int {

	if len(nums) <= 0 {
		return 0
	}
	dp := make(map[int]int)
	// dp[0] = 1

	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
func combinationSum42(nums []int, target int) int {

	dp := make(map[int]int)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
