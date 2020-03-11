package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println("res is ", res)
}

/*
4. 寻找两个有序数组的中位数
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0
示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return find3(nums1, nums2)
}

// 暴力法，先合并排序后，然后取数 不可取
func find1(nums1, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1) == 1 {
		return float64(nums1[0])
	}
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	} else {
		return (float64(nums1[len(nums1)/2]) + float64(nums1[len(nums1)/2-1])) * 0.5
	}
}

// 并不需要真正的合并，我们只是要找到合并后length1+length2的中位数，也就是（length1+length2）/2的数据，
// 因为可能是偶数，所以要保存前一个数
func find2(nums1, nums2 []int) float64 {
	left, right := math.MinInt32, math.MinInt32
	aStart, bStart := 0, 0
	length := len(nums1) + len(nums2)
	scanned := 0
	for aStart < len(nums1) || bStart < len(nums2) {
		left = right
		// 如果nums1,和nums2都还有数可以比较
		if aStart < len(nums1) && bStart < len(nums2) {
			if nums1[aStart] <= nums2[bStart] {
				right = nums1[aStart]
				aStart++
			} else if nums1[aStart] > nums2[bStart] {
				right = nums2[bStart]
				bStart++
			}
			//如果nums2已比较完
		} else if aStart < len(nums1) && bStart == len(nums2) {
			right = nums1[aStart]
			aStart++
			//如果nums1已比较完
		} else if aStart == len(nums1) && bStart < len(nums2) {
			right = nums2[bStart]
			bStart++
		}
		scanned++
		if scanned >= length/2+1 {
			break
		}
	}
	// 说明已扫描完成了
	if length%2 == 0 {
		return (float64(left) + float64(right)) / 2
	} else {
		return float64(right)
	}

}

func find3(nums1, nums2 []int) float64 {
	left := (len(nums1) + len(nums2) + 1) / 2
	right := (len(nums1) + len(nums2) + 2) / 2
	return float64(getK(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, left)+
		getK(nums1, 0, len(nums1)-1, nums2, 0, len(nums2)-1, right)) * 0.5
}

func getK(nums1 []int, start1, end1 int, nums2 []int, start2, end2, k int) float64 {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 < len2 {
		return getK(nums2, start2, end2, nums1, start1, end1, k)
	}
	if k == 1 {
		return math.Min(float64(nums1[start1]), float64(nums2[start2]))
	}

	i := start1 + int(math.Min(float64(len1), float64(k/2))-1)
	j := start2 + int(math.Min(float64(len2), float64(k/2))-1)
	if nums1[i] > nums2[j] {
		return getK(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getK(nums1, i+1, end1, nums2, j, end2, k-(i-start1+1))
	}
}
