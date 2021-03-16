package main

import (
	"fmt"
	"math"
	. "outback/leetcode/common"
)

func main() {
	res := maximumGap([]int{1, 1, 1, 1, 1, 5, 5, 5, 5, 5})
	fmt.Println(res)
}

/*
 给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
 如果数组元素个数小于 2，则返回 0。
 示例 1:
 输入: [3,6,9,1]
 输出: 3
 解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
 示例 2:
 输入: [10]
 输出: 0
 解释: 数组元素个数小于 2，因此返回 0。
 说明:
 你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
 请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
*/
// 是排序之后相临，而不是原数组相临

// 桶排序的思想
func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range nums {
		min, max = Min(min, n), Max(max, n)
	}
	//
	if min == max {
		return 0
	}
	// 分配桶的长度和个数是桶排序的关键
	// 在 n 个数下，形成的两两相邻区间是 n - 1 个，比如 [2,4,6,8] 这里
	// 有 4 个数，但是只有 3 个区间，[2,4], [4,6], [6,8]
	// 因此，桶长度 = 区间总长度 / 区间总个数 = (max - min) / (nums.length - 1)
	// int bucketSize = Math.max(1, (max - min) / (nums.length - 1));

	// 上面得到了桶的长度，我们就可以以此来确定桶的个数
	// 桶个数 = 区间长度 / 桶长度
	// 这里考虑到实现的方便，多加了一个桶，为什么？
	// 还是举上面的例子，[2,4,6,8], 桶的长度 = (8 - 2) / (4 - 1) = 2
	//                           桶的个数 = (8 - 2) / 2 = 3
	// 已知一个元素，需要定位到桶的时候，一般是 (当前元素 - 最小值) / 桶长度
	// 这里其实利用了整数除不尽向下取整的性质
	// 但是上面的例子，如果当前元素是 8 的话 (8 - 2) / 2 = 3，对应到 3 号桶
	//              如果当前元素是 2 的话 (2 - 2) / 2 = 0，对应到 0 号桶
	// 你会发现我们有 0,1,2,3 号桶，实际用到的桶是 4 个，而不是 3 个
	// 透过例子应该很好理解，但是如果要说根本原因，其实是开闭区间的问题
	// 这里其实 0,1,2 号桶对应的区间是 [2,4),[4,6),[6,8)
	// 那 8 怎么办？多加一个桶呗，3 号桶对应区间 [8,10)
	bucketSize := Max(1, (max-min)/(len(nums)-1)) // 桶间距至少为1, 桶里也可能有最大的间距
	// 题目中说了非负
	buckets := make([]*bucket, (max-min)/bucketSize+1)
	for _, n := range nums {
		loc := (n - min) / bucketSize
		if buckets[loc] == nil {
			buckets[loc] = &bucket{
				max: math.MinInt64,
				min: math.MaxInt64,
			}
		}
		buckets[loc].min = Min(buckets[loc].min, n)
		buckets[loc].max = Max(buckets[loc].max, n)
	}
	// 找值了
	ans := math.MinInt64
	preMax := math.MaxInt64
	for i := range buckets {
		if buckets[i] != nil {
			if preMax != math.MaxInt64 {
				ans = Max(ans, buckets[i].min-preMax)
			}
			preMax = buckets[i].max
			ans = Max(ans, buckets[i].max-buckets[i].min)
		}
	}
	return ans
}

type bucket struct {
	max int
	min int
}
