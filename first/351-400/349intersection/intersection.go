package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	res := intersection(nums1, nums2)
	fmt.Println("res is ", res)
}

/*
给定两个数组，编写一个函数来计算它们的交集。
示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
*/

// 使用两个Map
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	for _, num := range nums1 {
		map1[num] += 1
	}
	for _, num := range nums2 {
		map2[num] += 1
	}
	for k, _ := range map1 {
		r := map2[k]
		if r > 0 {
			//v := int(math.Min(float64(v), float64(r)))
			//for i := 1; i <= v; i++ {
			res = append(res, k)
		}
		//}
	}
	return res
}
