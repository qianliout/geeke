package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := triangleNumber([]int{2, 2, 3, 4})
	fmt.Println("ans is ", ans)
}

/*
给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
示例 1:
输入: [2,2,3,4]
输出: 3
解释:
有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3
*/

// 这种方法能得到结果，但是会超时
func triangleNumber(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if verification(nums, i, j, k) {
					ans++
				}
			}
		}
	}
	return ans
}

func verification(nums []int, i, j, k int) bool {
	if nums[i]+nums[j] <= nums[k] || nums[k]-nums[i] >= nums[j] || nums[k]-nums[j] >= nums[i] {
		return false
	}
	return true
}
