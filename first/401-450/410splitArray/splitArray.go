package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common"
)

func main() {
	nums := []int{7, 2, 5, 10, 8}
	res := splitArrary2(nums, 2)
	fmt.Println("res is ", res)
}

/*
给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
注意:
数组长度 n 满足以下条件:
1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
示例:
输入:
nums = [7,2,5,10,8]
m = 2
输出:
18
解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
*/
func splitArray(nums []int, m int) int {
	if len(nums) == 0 || m > len(nums) {
		return 0
	}

	max, sum := 0, 0
	for _, n := range nums {
		max = Max(max, n)
		sum += n
	}
	if len(nums) == m {
		return max
	}

	// 开始二分
	left, right := max, sum

	for left < right {
		mid := left + (right-left)/2
		// 这里一定要注意，当s==m时不能返回，意思就是，比如mid=21,此时也只能分两次，但他不是最小的值，所以还向左走，找最小值
		s := split(nums, mid)
		if s > m {
			// 如果分割数太多，说明「子数组各自的和的最大值」太小，此时需要将「子数组各自的和的最大值」调大
			// 下一轮搜索的区间是 [mid + 1, right]
			left = mid + 1
		} else {
			// 下一轮搜索的区间是上一轮的反面区间 [left, mid]
			right = mid
		}
	}
	return left
}

// 这个函数的意思是：当每个分组的值都不大于maxinter时，最少要分多少个组
func split(nums []int, maxinter int) int {
	s := 1
	curr := 0
	for _, n := range nums {
		if curr+n > maxinter {
			s++
			curr = 0
		}
		curr += n
	}
	return s
}

func splitArrary2(nums []int, m int) int {
	if len(nums) == 0 || m > len(nums) {
		return 0
	}

	all := 0
	max := math.MinInt64
	for _, n := range nums {
		all += n
		if n > max {
			max = n
		}
	}
	left, right := max, all

	// 开始二分
	for left <= right {
		mid := left + (right-left)/2
		n := split(nums, mid)
		if n == m {
			// 先不着急返回，锁定右边界，再继续找找，看能不能找到更小的
			fmt.Println("mid is ", mid)
			right = mid - 1
		} else if n > m {
			left = mid + 1
		} else if n < m {
			right = mid - 1
		}
	}
	return left
}
