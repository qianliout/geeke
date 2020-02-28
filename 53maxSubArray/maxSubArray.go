package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArray(nums)
	fmt.Println("max is ", max)
}

// 因为是连续，所以可以使用贪心算法
func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = int(math.Max(float64(nums[i]), float64(currSum+nums[i])))
		maxSum = int(math.Max(float64(maxSum), float64(currSum)))
	}
	return maxSum
}
