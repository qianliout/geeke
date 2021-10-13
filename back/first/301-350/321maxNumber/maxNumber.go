package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	// nums1 := []int{3, 4, 6, 5}
	// nums2 := []int{9, 1, 2, 5, 8, 3}
	nums1 := []int{2, 5, 6, 4, 4, 0}
	nums2 := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}
	res := maxNumber(nums1, nums2, 15)
	fmt.Println(res)
}

/*
给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n)
个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
说明: 请尽可能地优化你算法的时间和空间复杂度。
示例 1:
输入:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
输出:
[9, 8, 6, 5, 3]
示例 2:
输入:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
输出:
[6, 7, 6, 0, 4]
示例 3:
输入:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
输出:
[9, 8, 9]
*/

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, 0)
	
	for i := 0; i <= k; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			r := merge(pixMax(nums1, i), pixMax(nums2, k-i))
			if compare(r, res) {
				res = r
			}
		}
	}
	return res
}

func pixMax(nums []int, k int) []int {
	res := make([]int, 0)
	drop := len(nums) - k
	for _, num := range nums {
		for drop > 0 && len(res) > 0 && res[len(res)-1] < num {
			drop--
			res = res[:len(res)-1]
		}
		res = append(res, num)
	}
	fmt.Println("pix max is ", res[:k])
	return res[:k]
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, 0)
	for len(nums1) > 0 && len(nums2) > 0 {
		// 千万不能这样写 if nums1[0]>nums2[0]  [6,7],[6,0,4]
		// [2 5 6 4 4 0] . [7 3 8 0 6 5 7 6 2]
		// 这里就是这道题的难点
		if compare(nums1, nums2) {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	if len(nums1) > 0 {
		res = append(res, nums1...)
	}
	if len(nums2) > 0 {
		res = append(res, nums2...)
	}
	fmt.Println("merge is ", res)
	return res
}

func compare(nums1, nums2 []int) bool {
	for i := 0; i < common.Min(len(nums1), len(nums2)); i++ {
		if nums1[i] > nums2[i] {
			return true
		} else if nums1[i] < nums2[i] {
			return false
		}
	}
	if len(nums1) > len(nums2) {
		return true
	}
	return false
}
