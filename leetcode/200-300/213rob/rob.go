package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{2, 3, 2}
	num := rob(nums)
	fmt.Println("rob ", num)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(help(nums[:len(nums)-1]), help(nums[1:]))
}

func help(nums []int) int {
	rob, notRob := make([]int, len(nums)), make([]int, len(nums))
	rob[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		rob[i] = Max(notRob[i-1]+nums[i], nums[i])
		notRob[i] = Max(notRob[i-1], rob[i-1])
	}

	return Max(rob[len(nums)-1], notRob[len(nums)-1])
}
