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

// 桶的思想
// 这里要防止移出
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 || len(nums) <= 0 {
		return false
	}
	bucket := make(map[int]int)

	for i := range nums {
		// 这一步就是这道题的关键
		// 桶的大小设成t+1更加方便
		nth := int(math.Floor(float64(nums[i]) / float64(t+1))) // 放入那个桶
		if _, exit := bucket[nth]; exit {
			return true // 如果放入之前桶中已经有了一个了,说明新放入的这个元素加上这个,相差就不会超过t
		}
		if _, exit := bucket[nth-1]; exit && int(math.Abs(float64(bucket[nth-1]-nums[i]))) <= t {
			return true
		}
		if _, exit := bucket[nth+1]; exit && int(math.Abs(float64(bucket[nth+1]-nums[i]))) <= t {
			return true
		}
		bucket[nth] = nums[i]
		// fmt.Println("bucket is ", bucket)
		// 把距离超过k之后的数据删除,这里删除的目的主要是为了每一个存在的判断,如果不删除,上面就要使用循环
		if i >= k {
			delete(bucket, nums[i-k]/(t+1))
		}
	}
	return false
}
