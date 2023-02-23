package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	ans := findMaximumXOR(nums)
	fmt.Println("ans is ", ans)
}

/*
421. 数组中两个数的最大异或值
给定一个非空数组，数组中元素为 a0, a1, a2, … , an-1，其中 0 ≤ ai < 231 。
找到 ai 和aj 最大的异或 (XOR) 运算结果，其中0 ≤ i,  j < n 。
你能在O(n)的时间解决这个问题吗？
示例:
输入: [3, 10, 5, 25, 2, 8]
输出: 28
解释: 最大的结果是 5 ^ 25 = 28.
*/
func findMaximumXOR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return find(nums)
}

// https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/solution/li-yong-yi-huo-yun-suan-de-xing-zhi-tan-xin-suan-f/
func find(nums []int) int {
	res := 0
	mask := 0
	for i := 30; i >= 0; i-- {
		mask = mask | (1 << i)
		set := make(map[int]interface{})
		for _, num := range nums {
			set[mask&num] = nil
		}
		temp := res | (1 << i)
		for key := range set {
			if _, ok := set[key^temp]; ok {
				res = temp
				break
			}
		}
	}
	return res
}

func Max(nums []int) int {
	m := math.MinInt64
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
