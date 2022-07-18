package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	ans := minSubArrayLen(7, nums)
	fmt.Println(ans)
}

func minSubArrayLen(target int, nums []int) int {
	left, right, ans, tmp := 0, 0, math.MaxInt64, 0

	for left < len(nums) && right < len(nums) {
		tmp += nums[right]
		right++
		for tmp >= target {
			ans = Min(ans, right-left)
			tmp -= nums[left]
			left++
		}
	}
	if ans == math.MinInt64 {
		return 0
	}
	return ans
}
