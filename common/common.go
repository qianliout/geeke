package common

import (
	"fmt"
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
func BitLeft(nums *[]int, tar int) int {
	if len(*nums) == 0 {
		*nums = append(*nums, tar)
		return 0
	}
	length := len(*nums)
	left, right := 0, length-1
	for left < right {
		mid := left + (right-left)/2
		if (*nums)[mid] < tar {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//  检查最后
	if (*nums)[left] < tar {
		*nums = append(*nums, tar)
		return left
	}

	fmt.Println("num is ", *nums, left, tar)
	r := (*nums)[left:]
	*nums = append(append((*nums)[:left+1], tar), r...)

	return left
}
