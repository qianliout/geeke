package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{1, 2, 3, 1, 4, 10}
	ans := rob(nums)
	fmt.Println(ans)
}

func rob(nums []int) int {
	hold, notHold := make([]int, len(nums)), make([]int, len(nums))
	hold[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 下手
		hold[i] = Max(notHold[i-1] + nums[i])
		// 不下手
		notHold[i] = Max(notHold[i-1], hold[i-1])
	}
	return Max(Max(hold...), Max(notHold...))
}
