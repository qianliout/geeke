package main

import (
	"math"
)

func main() {

}

/*
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
示例 1:
输入: [3,2,3]
输出: [3]
示例 2:
输入: [1,1,1,3,3,2,2,2]
输出: [1,2]
*/

func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if len(nums) <= 0 {
		return []int{}
	}
	first, sencond := nums[0], nums[0]
	firstN, sencondN := 0, 0
	// 统计阶段
	for i := 0; i < len(nums); i++ {
		if nums[i] == first {
			firstN++
			continue
		}
		if nums[i] == sencond {
			sencondN++
			continue
		}
		if firstN == 0 {
			first = nums[i]
			firstN++
			continue
		}
		if sencondN == 0 {
			sencond = nums[i]
			sencondN++
			continue
		}
		firstN--
		sencondN--
	}
	// 计数阶段
	firstN, sencondN = 0, 0
	for _, i := range nums {
		if i == first {
			firstN++
		}
		if i == sencond {
			sencondN++
		}
	}
	if firstN > int(math.Ceil(float64(len(nums)/3))) {
		res = append(res, first)
	}

	if sencondN > int(math.Ceil(float64(len(nums)/3))) && first != sencond {
		res = append(res, sencond)
	}
	return res
}
