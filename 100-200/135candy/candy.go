package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	can := candy([]int{1, 0, 2})
	fmt.Println("candy is ", can)
}

/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
*/
func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	left, right := make([]int, len(ratings)), make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i], right[i] = 1, 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(right) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := 0
	for i := 0; i < len(ratings); i++ {
		ans += Max(left[i], right[i])
	}
	return ans
}
