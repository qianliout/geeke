package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lis := lengthOfLIS(nums)
	fmt.Println("list ", lis)
}

func lengthOfLIS(nums []int) int {
	ans := 0
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		ans = Max(ans, dp[i])
	}
	return ans
}
