package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	ans := productExceptSelf(nums)
	fmt.Println("ans is ", ans)
}

func productExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)+1), make([]int, len(nums)+1)
	left[0], right[len(right)-1] = 1, 1
	res := make([]int, len(nums))
	// 左边
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] * nums[i]
	}
	fmt.Println("left is ", left)
	// 右边
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}
	fmt.Println("right is ", right)

	for i := 0; i < len(nums); i++ {
		res[i] = left[i] * right[i+1]
	}
	return res
}
