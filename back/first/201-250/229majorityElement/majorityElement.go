package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 2}
	res := majorityElement2(nums)
	fmt.Println("res is ", res)
}

/*
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
示例 1:
输入: [3,2,3]
输出: [3]
示例 2:
输入: [1,1,1,3,3,2,2,2]
输出: [1,2]
*/

// 这道题使用map是很容易的,但是空间复杂度不符合要求,所以使用moer投票法
// 以后这种求众数的题都是这样做
//超过 ⌊ n/3 ⌋ 说明最后只有两个元素1
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	cand1, cand2, count1, count2 := nums[0], nums[0], 0, 0

	// 投票阶段
	for _, num := range nums {
		if cand1 == num {
			count1++
			continue
		}
		if cand2 == num {
			count2++
			continue
		}
		if count1 == 0 {
			cand1 = num
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = num
			count2++
			continue
		}

		count1--
		count2--
	}

	// 计数阶段

	count1, count2 = 0, 0

	for _, num := range nums {
		if num == cand1 {
			count1++
		} else if num == cand2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}

// 彻底掌握摩尔投票法

func majorityElement2(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	first, second, f, s := nums[0], nums[0], 0, 0
	// 抵消阶段
	for _, nu := range nums {
		if nu == first {
			f++
			continue
		}
		if nu == second {
			s++
			continue
		}
		if f == 0 {
			first = nu
			f = 1
			continue
		}
		if s == 0 {
			second = nu
			s = 1
			continue
		}
		f--
		s--
	}
	// 计数阶段
	f, s = 0, 0
	for _, nu := range nums {
		if nu == first {
			f++
		}
		if nu == second {
			s++
		}
	}

	if f > int(math.Ceil(float64(len(nums)/3))) {
		res = append(res, first)
	}

	if s > int(math.Ceil(float64(len(nums)/3))) && first != second {
		res = append(res, second)
	}
	return res
}
