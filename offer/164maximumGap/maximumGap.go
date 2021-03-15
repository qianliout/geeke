package main

import (
	"math"
	. "outback/leetcode/common"
)

func main() {

}

/*
 给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
 如果数组元素个数小于 2，则返回 0。
 示例 1:
 输入: [3,6,9,1]
 输出: 3
 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
 示例 2:
 输入: [10]
 输出: 0
 解释: 数组元素个数小于 2，因此返回 0。
 说明:
 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
*/
// 是排序之后相临，而不是原数组相临

// 桶排序的思想
func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range nums {
		min, max = Min(min, n), Max(max, n)
	}
	//
	if min == max {
		return 0
	}
	// 确定桶的个数
	bucketSize := int(math.Ceil(math.Max(1, float64(max-min)) / float64(len(nums)-1)))

}
