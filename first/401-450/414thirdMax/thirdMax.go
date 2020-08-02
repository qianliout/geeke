package main

import (
	"fmt"
	"math"
)

func main() {
	res := thirdMax([]int{3, 2, 1})
	fmt.Println("res is ", res)
}

/*
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
示例 1:
输入: [3, 2, 1]
输出: 1
解释: 第三大的数是 1.
示例 2:
输入: [1, 2]
输出: 2
解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:
输入: [2, 2, 3, 1]
输出: 1
解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。
*/

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	first, second, third := nums[0], math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n > first {
			third = second
			second = first
			first = n
		} else if n < first && n > second {
			third = second
			second = n
		} else if n < second && n > third {
			third = n
		}
	}
	if third == math.MinInt64 {
		return first
	}
	return third
}
