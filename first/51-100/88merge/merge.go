package main

import (
	"fmt"
)

func main() {
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)

}

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明:
    初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
    你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/

// 归并排序的思想,使用双指
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p >= 0 {
		// 把nums1移动到可以接收下一个nums2为止
		if (p2 < 0 && p1 >= 0) || (p1 >= 0 && nums1[p1] >= nums2[p2]) {
			nums1[p] = nums1[p1]
			p1--
		} else {
			// 放nums2
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	fmt.Println(nums1)
}
