package main

import (
	"sort"
)

func main() {
	nums := []int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3}
	wiggleSort(nums)
}

/*
给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
你可以假设所有输入数组都可以得到满足题目要求的结果。
*/
func wiggleSort(nums []int) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	// 这一步是重点，当数组长度是奇数时，保证前面比后面多
	mid := (len(nums) + 1) / 2

	nums1, nums2 := tmp[:mid], tmp[mid:]
	// 反序
	for i := 0; i < len(nums1)/2; i++ {
		nums1[len(nums1)-i-1], nums1[i] = nums1[i], nums1[len(nums1)-i-1]
	}

	for i := 0; i < len(nums2)/2; i++ {
		nums2[len(nums2)-i-1], nums2[i] = nums2[i], nums2[len(nums2)-i-1]
	}

	i := 0
	for len(nums1) > 0 && len(nums2) > 0 {
		nums[i] = nums1[0]
		i++
		nums[i] = nums2[0]
		i++
		nums1, nums2 = nums1[1:], nums2[1:]
	}
	for j := range nums1 {
		nums[i] = nums1[j]
		i++
	}
	for j := range nums2 {
		nums[i] = nums2[j]
		i++
	}
}
