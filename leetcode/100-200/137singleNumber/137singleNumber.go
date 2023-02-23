package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}

	fmt.Println(singleNumber(nums))
}

/*
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
*/
func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		count := 0
		for _, num := range nums {
			count += (num >> i) & 1
		}
		if count%3 == 1 {
			ans += 1 << i
		}
	}
	return ans
}
