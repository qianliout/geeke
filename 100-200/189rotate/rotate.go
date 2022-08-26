package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

/*
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
*/
func rotate2(nums []int, k int) {
	k = k % len(nums)
	if k == 0 || len(nums) <= 1 {
		return
	}
	tmm := append([]int{}, nums[len(nums)-k:]...)
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmm[i]
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	rotateAll(nums)
	rotateAll(nums[:k])
	rotateAll(nums[k:])
}

func rotateAll(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
