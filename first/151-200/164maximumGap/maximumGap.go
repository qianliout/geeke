package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// nums := []int{15252, 16764, 27963, 7817, 26155, 20757, 3478, 22602, 20404, 6739, 16790, 10588, 16521, 6644, 20880, 15632, 27078, 25463, 20124, 15728, 30042, 16604, 17223, 4388, 23646, 32683, 23688, 12439, 30630, 3895, 7926, 22101, 32406, 21540, 31799, 3768, 26679, 21799, 23740}
	nums := []int{3, 6, 9, 1}
	res := maximumGap3(nums)
	fmt.Println("res is ", res) // 2901
}

/*
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
如果数组元素个数小于 2，则返回 0。
示例 1:
输入: [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。
示例 2:
输入: [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。
说明:
    你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
    请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
链接：https://leetcode-cn.com/problems/maximum-gap
*/

// 方法一,暴力排序  (也可以使用基数排序的方法，但都是排序)
func maximumGap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	res := math.MinInt64
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] > res {
			res = nums[i] - nums[i-1]
		}
	}
	return res
}

// 这里并不是要实质的全部排序，只是求一个最大差值，所以可以使用桶排序的思想
/*
划定了箱子范围后，我们其实很容易把数字放到箱子中，通过 (nums[i] - min) / interval 即可得到当前数字应该放到的箱子编号。
那么最主要的问题其实就是怎么去确定 interval。
interval 过小的话，需要更多的箱子去存储，很费空间，此外箱子增多了，比较的次数也会增多，不划算。
interval 过大的话，箱子内部的数字可能产生题目要求的最大 gap，所以肯定不行。
所以我们要找到那个保证箱子内部的数字不会产生最大 gap，并且尽量大的 interval。
继续看上边的例子，0 3 4 6 23 28 29 33 38，数组中的最小值 0 和最大值 38 ,并没有参与到 interval 的计算中，所以它俩可以不放到箱子中，还剩下 n - 2 个数字。
像上边的例子，如果我们保证至少有一个空箱子，那么我们就可以断言，箱子内部一定不会产生最大 gap。
因为在我们的某次计算中，会跳过一个空箱子，那么得到的 gap 一定会大于 interval，而箱子中的数字最大的 gap 是 interval - 1。
接下来的问题，怎么保证至少有一个空箱子呢？
鸽巢原理的变形，有 n - 2 个数字，如果箱子数多于 n - 2 ，那么一定会出现空箱子。总范围是 max - min，
那么 interval = (max - min) / 箱子数，为了使得 interval 尽量大，箱子数取最小即可，也就是 n - 1。
所以 interval = (max - min) / n - 1 。这里如果除不尽的话，我们 interval 可以向上取整。
因为我们给定的数字都是整数，这里向上取整的话对于最大 gap 是没有影响的。比如原来范围是 [0,5.5)，
那么内部产生的最大 gap 是 5 - 0 = 5。现在向上取整，范围变成[0,6)，但是内部产生的最大 gap 依旧是 5 - 0 = 5。
链接：https://leetcode-cn.com/problems/maximum-gap/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--39/
*/
func maximumGap2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	min := math.MaxInt64
	max := math.MinInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	if max == min {
		return 0
	}
	bucketSize := int(math.Ceil(math.Max(1, float64(max-min)) / float64(len(nums)-1)))
	fmt.Println("buckSize", bucketSize)
	fmt.Println("min max ", min, max)
	maxMap := make(map[int]int)
	minMap := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		maxMap[i] = math.MinInt64
		minMap[i] = math.MaxInt64
	}
	// 入桶,并不是真正的放入桶中,只是记录,这个桶中的最大值,和最小值
	for _, num := range nums {
		if num == max || num == min {
			continue
		}
		index := (num - min) / bucketSize
		if maxMap[index] < num {
			maxMap[index] = num
		}
		if minMap[index] > num {
			minMap[index] = num
		}
	}
	res := math.MinInt64
	preMax := min
	for i := 0; i < len(nums)-1; i++ {
		if maxMap[i] == math.MinInt64 {
			continue
		}
		if minMap[i]-preMax > res {
			res = minMap[i] - preMax
		}
		preMax = maxMap[i]
	}

	// 最大值可能处于边界，不在箱子中，需要单独考虑
	if max-preMax > res {
		res = max - preMax
	}
	// fmt.Println(":max ", maxMap)
	// fmt.Println(":min ", minMap)

	return res
}

// 这种方法，用len(num)+1个桶，在常规情况下可以得到正确的解，但是在一些情况下是得不到解的
// 1,[1,10000] 2,[1,1,1,1,1,1]
func maximumGap3(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	max, min := math.MinInt64, math.MaxInt64
	for _, n := range nums {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	// 这里设定有len(nums)+1个桶

	gap := int(math.Ceil(float64(max-min) / (float64(len(nums) + 1))))
	// gap := 2
	// 不是真正放入桶中，只是记录各个桶中的最大值，最小值
	maxMap := make(map[int]int)
	minMap := make(map[int]int)
	for i := 0; i < len(nums)+1; i++ {
		maxMap[i] = math.MinInt64
		minMap[i] = math.MaxInt64
	}

	for _, n := range nums {
		index := (n - min) / gap
		if maxMap[index] < n {
			maxMap[index] = n
		}
		if minMap[index] > n {
			minMap[index] = n
		}
	}
	// 找到最大的间隔，最大的间隔就在两个桶之间
	res, pre := 0, min
	for i := 0; i < len(nums)+1; i++ {
		if maxMap[i] == math.MinInt64 {
			continue // 说明这个桶没有值
		}
		if minMap[i]-pre > res {
			res = minMap[i] - pre
		}
		pre = maxMap[i]
	}

	return res
}
