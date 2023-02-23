package main

import (
	"fmt"
)

func main() {
	nums1 := []int{0, 1, 2, 3}
	nums2 := []int{1, 1, 2, 3}
	res := dichotomy(nums1, nums2)
	fmt.Println(res)
}

/*
	给定两个数组，编写一个函数来计算它们的交集。

示例 1:
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

	输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
	我们可以不考虑输出结果的顺序。

进阶:

	如果给定的数组已经排好序呢？你将如何优化你的算法？
	如果 nums1 的大小比 nums2 小很多，哪种方法更优？
	如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	return useMap(nums1, nums2)
}

// 使用map ,可以进行优化，只存一个map,然后遍历另一个，少一次循环
func useMap(nums1, nums2 []int) []int {
	res := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return res
	}
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, num := range nums1 {
		map1[num]++
	}
	for _, num := range nums2 {
		map2[num]++
	}

	if len(nums1) <= len(nums2) {
		for _, num := range nums1 {
			if v, exit := map2[num]; exit && v > 0 {
				map2[num]--
				res = append(res, num)
			}
		}
	} else {
		for _, num := range nums2 {
			if v, exit := map1[num]; exit && v > 0 {
				map1[num]--
				res = append(res, num)
			}
		}
	}
	return res
}

// 如果给定的数组已经排好序呢？你将如何优化你的算法,可以使用双指针的方式
func dichotomy(nums1, nums2 []int) []int {
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			fmt.Println("i is ,j is ,k is ", i, j, k)
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
