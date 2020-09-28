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

func Abs(n1, n2 int) int {
	return int(math.Abs(float64(n1 - n2)))
}

// 二分法,找到插入左边的位置,并插入左边,和python的bitset_left一样
func FindSmallIdxAndInsert(sorted *[]int, target int) int {

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

// 二分法,找到插入左边的位置,并插入左边,和python的bitset_left一样
func FindSmallIdx(sorted *[]int, target int) int {

	if len(*sorted) == 0 {
		return 0
	}
	start := 0
	end := len(*sorted) - 1

	if (*sorted)[end] < target {
		return end + 1
	}
	if (*sorted)[start] > target {
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
	return start
}

func Insert(a *[]int, c int, i int) []int {
	*a = append((*a)[:i], append([]int{c}, (*a)[i:]...)...)
	return *a
}

func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}
