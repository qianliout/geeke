package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	duplicate := containsNearbyAlmostDuplicate(nums, 2, 3)
	fmt.Println("dup", duplicate)
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	for i := range nums {
		for j := i - 1; i-j <= k; j-- {
			if j < 0 {
				break
			}
			if Abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

// 使用桶排序的思想
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	bucketSize := t + 1
	bucket := make(map[int]int)

	for i := range nums {
		nth := int(math.Floor(float64(nums[i]) / float64(bucketSize)))
		if _, ok := bucket[nth]; ok {
			return true
		}
		if _, ok := bucket[nth-1]; ok && Abs(bucket[nth-1]-nums[i]) <= t {
			return true
		}
		if _, ok := bucket[nth+1]; ok && Abs(bucket[nth+1]-nums[i]) <= t {
			return true
		}
		bucket[nth] = nums[i]
		if i >= k {
			delete(bucket, int(math.Floor(float64(nums[i-k])/float64(bucketSize))))
		}
	}
	return false
}
