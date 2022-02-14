package main

import (
	"fmt"
	"math"

	"qianliout/leetcode/common/utils"
)

func main() {
	// [2147483647,-1,2147483647]
	// 1
	// 2147483647

	res := containsNearbyAlmostDuplicate([]int{2147483647, -1, 2147483647}, 1, 2147483647)
	fmt.Println(res)

	fmt.Println(-1 / float64(2147483648))
	fmt.Println(math.Floor(-1 / float64(2147483648)))
}

/*
在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的绝对值也小于等于 ķ 。
如果存在则返回 true，不存在返回 false。
示例 1:
输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:
*/

// 桶排序的思想,桶的大小，有多少个桶，这个还不太会,要复习
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucket := make(map[int]int)
	for i := range nums {
		con1 := int(math.Floor(float64(nums[i]) / float64(t+1)))
		fmt.Println("conn1 ", con1, nums[i], float64(t+1))
		con := nums[i] / (t + 1)
		fmt.Println("conn2 ", con, nums[i], t+1)

		if _, ok := bucket[con]; ok {
			return true
		}
		if _, ok := bucket[con-1]; ok && utils.AbsSub(bucket[con-1], nums[i]) <= t {
			return true
		}
		if _, ok := bucket[con+1]; ok && utils.AbsSub(bucket[con+1], nums[i]) <= t {
			return true
		}
		bucket[con] = nums[i]
		if i >= k {
			delete(bucket, nums[i-k]/(t+1))
		}
	}
	return false
}
