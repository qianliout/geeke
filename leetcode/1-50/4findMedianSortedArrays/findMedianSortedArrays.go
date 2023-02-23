package main

import (
	"sort"
)

func main() {

}

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。
*/
// 方法一，排序
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1) == 0 {
		return 0
	}
	if len(nums1)%2 != 0 {
		return float64(nums1[len(nums1)/2])
	}
	return (float64(nums1[len(nums1)/2]) + float64(nums1[len(nums1)/2-1])) / 2
}

// 模拟扫描,这里的数组是正序列数组
// 1 <= m + n <= 2000 // 也就是说只能是一个为空
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	leftStart, rightStart := 0, 0
	left, right := -1, -1 //  因为是正序列数组
	for i := 0; i <= (len(nums1)+len(nums2))/2; i++ {
		left = right
		// 这里的条件判断是这个题的难点
		if leftStart < len(nums1) && (rightStart >= len(nums2) || nums1[leftStart] < nums2[rightStart]) {
			right = nums1[leftStart]
			leftStart++
		} else {
			right = nums2[rightStart]
			rightStart++
		}
	}
	if (len(nums1)+len(nums2))&1 == 0 {
		return (float64(left) + float64(right)) / 2
	}
	return float64(right)
}
