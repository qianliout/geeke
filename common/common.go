package common

import (
	"math"
)

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 二分法,找到插入左边的位置,并插入左边
func FindSmallIdx(sorted *[]int, target int) int {

	if len(*sorted) == 0 {
		Insert(sorted, target, 0)
		return 0
	}
	start := 0
	end := len(*sorted) - 1

	if (*sorted)[end] < target {
		Insert(sorted, target, end+1)
		return end + 1
	}
	if (*sorted)[start] > target {
		Insert(sorted, target, 0)
		return 0
	}

	for start < end {
		mid := start + (end-start)/2
		if (*sorted)[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	// 插入数据
	Insert(sorted, target, start)
	return start
}

func Insert(a *[]int, c int, i int) []int {
	*a = append((*a)[:i], append([]int{c}, (*a)[i:]...)...)
	return *a
}
