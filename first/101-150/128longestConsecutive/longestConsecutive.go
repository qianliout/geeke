package main

import (
	"fmt"
	"math"
	"sort"

	"outback/leetcode/common/unionfind"
)

func main() {
	// nums := []int{1, 2, 0, 1}
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive3(nums)
	fmt.Println("res is ", res)
}

/*
给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。
示例:
输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
*/
// 方法一，排序会超过时间限制
func longestConsecutive(nums []int) int {

	sort.Ints(nums)

	if len(nums) <= 1 {
		return len(nums)
	}

	n := len(nums)
	ans := 1
	start := 1
	for start < n {
		tem := 1
		for start < n {
			if nums[start]-nums[start-1] == 1 {
				tem++
				start++
			} else if nums[start]-nums[start-1] == 0 {
				start++
			}
		}
		start++
		if tem > ans {
			ans = tem
		}
	}
	return ans
}
func longestConsecutive2(nums []int) int {
	max := math.MinInt64
	min := math.MaxInt64

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Println(min, max)
	uf := unionfind.NewUnionFind(max - min + 1)

	for _, num := range nums {
		if num > min {
			uf.Union(num-min, num-1-min)
		}
		if num < max {
			uf.Union(num-min, num+1-min)
		}
	}
	return max - min - uf.Count
}

// 使用map
func longestConsecutive3(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return len(nums)
	}

	m := make(map[int]int)

	max := math.MinInt64
	min := math.MaxInt64

	for _, num := range nums {
		m[num]++
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	ans := 0

	for _, num := range nums {
		if m[num-1] <= 0 {
			cruu := num
			this := 1
			for m[cruu+1] > 0 {
				cruu++
				this++
			}
			if this > ans {
				ans = this
			}
		}
	}
	return ans
}
