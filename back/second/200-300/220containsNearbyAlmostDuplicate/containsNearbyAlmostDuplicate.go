package main

import (
	"qianliout/leetcode/common/utils"
)

func main() {

}

/*
在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的绝对值也小于等于 ķ 。
如果存在则返回 true，不存在返回 false。
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

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if utils.AbsSub(nums[i], nums[j]) <= t && utils.AbsSub(i, j) <= k {
				return true
			}
		}
	}
	return false
}
