package main

import (
	"math"
	"sort"

	. "qianliout/leetcode/common/utils"
)

func main() {

}

func maximumGap(nums []int) int {

}

func useSort(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > ans {
			ans = nums[i] - nums[i-1]
		}
	}
	return ans
}

func bulk(nums []int) int {
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range nums {
		min, max = Min(min, n), Max(max, n)
	}
	return 0
}
