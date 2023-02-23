package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1, 2} // 这个就是个不好
	nextPermutation(num)
	fmt.Println("nums is ", num)

}

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须原地修改，只允许使用额外常数空间。
以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	ln := len(nums)
	flag := false

	// find start
	start := 0
	for i := ln - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			start = i - 1
			flag = true
			break
		}
	}
	// 找到下一次比start大的,但是是后面排列中最小的一个数
	if !flag {
		sort.Ints(nums)
		return
	}

	nexMax := start + 1
	for j := nexMax; j < ln; j++ {
		if nums[j] < nums[nexMax] && nums[j] > nums[start] {
			nexMax = j
		}
	}

	// 交换值
	nums[start], nums[nexMax] = nums[nexMax], nums[start]
	// 排列后面的值
	sort.Ints(nums[start+1:])
}
