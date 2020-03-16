package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-3, 3}
	fmt.Println(containsNearbyAlmostDuplicate(nums, 2, 4))
	fmt.Println(math.Floor(float64(-3) / float64(5)))
	fmt.Println(math.Floor(float64(3) / float64(5)))
}

/*
给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，并且 i 和 j 之间的差的绝对值最大为 ķ。
示例 1:
输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:
输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:
输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false
*/

// 这里要防止移出
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 || len(nums) <= 0 {
		return false
	}
	bucket := make(map[int]int)

	for i := range nums {
		// 这一步就是这道题的关键
		nth := int(math.Floor(float64(nums[i]) / float64(t+1)))
		if _, exit := bucket[nth]; exit {
			return true
		}
		if _, exit := bucket[nth-1]; exit && int(math.Abs(float64(bucket[nth-1]-nums[i]))) <= t {
			return true
		}
		if _, exit := bucket[nth+1]; exit && int(math.Abs(float64(bucket[nth+1]-nums[i]))) <= t {
			return true
		}
		bucket[nth] = nums[i]
		//fmt.Println("bucket is ", bucket)
		if i >= k {
			delete(bucket, nums[i-k]/(t+1))
		}
	}
	return false
}
