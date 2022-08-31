package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3}
	number := missingNumber(num)
	fmt.Println("res is ", number)
}

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
*/

func missingNumber(nums []int) int {
	diff := len(nums)
	for i, n := range nums {
		diff = diff ^ i ^ n
	}
	return diff
}
