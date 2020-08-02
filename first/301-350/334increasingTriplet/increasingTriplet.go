package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{0, 4, 2, 1, 0, -1, -3}
	triplet := increasingTriplet1(nums)
	fmt.Println("res is ", triplet)
}

/*
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。
数学表达式如下:
    如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
    使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。
示例 1:
输入: [1,2,3,4,5]
输出: true
示例 2:
输入: [5,4,3,2,1]
输出: false
*/

// 注意,这里没有说是连续子序列,以下解法是连续子序列
func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	first := 0
	length := len(nums)
	second := first + 1
	third := second + 1
	for first < length && second < length && third < length {
		if nums[second] <= nums[first] {
			first = second
			second = first + 1
			third = second + 1
			continue
		} else if nums[third] <= nums[second] {
			first = third
			second = third + 1
			third = second + 1
			continue
		} else {

			return true
		}
	}
	return false
}

func increasingTriplet1(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	first := nums[0]
	second := math.MaxInt64

	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] < first {
			first = nums[i]
			continue
		} else if nums[i] > first && nums[i] < second {
			second = nums[i]
			continue
		} else if nums[i] > second {
			fmt.Println(first, second, i)
			return true
		}
	}
	return false
}
