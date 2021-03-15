package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

}

/*
在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，且满足 i 和 j 的差的绝对值也小于等于 ķ 。
如果存在则返回 true，不存在返回 false。
示例 1:
输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:
*/

// 桶排序的思想,桶的大小，有多少个桶，这个还不太会,要复习
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
