package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	res := missingNumber(nums)
	fmt.Println("res is ", res)
}

/*
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
示例 1:
输入: [3,0,1]
输出: 2
示例 2:
输入: [9,6,4,2,3,5,7,0,1]
输出: 8
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?
链接：https://leetcode-cn.com/problems/missing-number
*/

// 这道题首先想到的是map，太简单，这里不写、
// 方法二，使用位运算
// 对一个数据进行两次异或运算就会得到原来的数
func missingNumber(nums []int) int {
	diff := 0
	for i := 0; i <= len(nums); i++ {
		diff = diff ^ i
	}

	for _, num := range nums {
		diff = diff ^ num
	}
	return diff
}
