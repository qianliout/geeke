package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0}
	nums2 := []int{2}
	merge(nums1, 0, nums2, 1)
	fmt.Println(nums1)
}

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明:
    初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/

// 归并排序的思想
func merge(nums1 []int, m int, nums2 []int, n int) {
	p, left, right := m+n-1, m-1, n-1
	for p >= 0 {
		if (left >= 0 && right < 0) || (left >= 0 && right >= 0 && nums1[left] >= nums1[right]) {
			nums1[p] = nums1[left]
			left--
		} else {
			nums1[p] = nums2[right]
			right--
		}
		p--
	}
}
