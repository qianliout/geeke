package main

import (
	"fmt"
)

func main() {
	//nums := []int{20, 2, 2, 2, 3, 3, 3}
	nums := []int{-1, -1, -1, -2}
	res := singleNumber(nums)
	fmt.Println("res is ", res)
}

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,3,2]
输出: 3
示例 2:
输入: [0,1,0,1,0,1,99]
输出: 99
链接：https://leetcode-cn.com/problems/single-number-ii
*/
func singleNumber(nums []int) int {
	return find2(nums)
}

// 也可以使用map，这里就不写了，太简单了

/*
https://leetcode-cn.com/problems/single-number-ii/solution/single-number-ii-mo-ni-san-jin-zhi-fa-by-jin407891/
*/
func find1(nums []int) int {
	ones, twos, threes := 0, 0, 0
	for _, num := range nums {
		twos |= ones & num
		ones ^= num
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}
	return ones
}

// 使用遍历的方式,这里就有计算机补码的方式,这里要特别考虑才行呢
// fixme 这里的写法有问题
func find2(nums []int) int {
	counts := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			counts[i] += num & 1
			num = num >> 1
		}
	}
	fmt.Println(counts)
	res := 0
	for i := 0; i < 32; i++ {
		res = res << 1
		res = res | counts[31-i]%3
	}

	if counts[31]%3 == 0 {
		return res
	}
	return -res
}
