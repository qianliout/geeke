package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 3, 3, 4}
	res := singleNumber(nums)
	fmt.Println("res is ", res)
}

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4
链接：https://leetcode-cn.com/problems/single-number
*/
func singleNumber(nums []int) int {
	return find2(nums)
}

// 按位操作
func find2(nums []int) int {
	res := 0

	for _, num := range nums {
		res = res ^ num
	}
	return res
}

// 使用map
func find1(nums []int) int {
	exit := make(map[int]int)
	for _, num := range nums {
		exit[num] += 1
	}

	for k, v := range exit {

		if v == 1 {
			return k
		}
	}
	return -1
}
