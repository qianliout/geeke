package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{3, 4, 6, 5}
	nums2 := []int{9, 1, 2, 5, 8, 3}
	number := maxNumber(nums1, nums2, 5)
	fmt.Println(number)
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, 0)

	for i := 0; i <= k; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			r := merge(findMax(nums1, i), findMax(nums2, k-i))
			if compare(r, ans) {
				ans = r
			}
		}
	}
	return ans
}

func findMax(nums []int, k int) []int {
	res := make([]int, 0)
	dropped := len(nums) - k

	for i := range nums {
		for dropped > 0 && len(res) > 0 && res[len(res)-1] < nums[i] {
			res = res[:len(res)-1]
			dropped--
		}
		res = append(res, nums[i])
	}
	return res[:k]
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, 0)
	for len(nums1) > 0 && len(nums2) > 0 {
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

	return res
}

func compare(nums1, nums2 []int) bool {

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			return true
		} else if nums1[i] < nums2[j] {
			return false
		} else {
			i, j = i+1, j+1
		}
	}
	return len(nums1) > len(nums2)
}

func compare2(nums1, nums2 []int) bool {

	for i := 0; i < Min(len(nums1), len(nums2)); i++ {
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

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
