package main

import (
	"fmt"
)

func main() {
	res := firstMissingPositive([]int{1, 2, 3, 4})
	fmt.Println("res is ", res)
}

/*
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
示例 1:
输入: [1,2,0]
输出: 3
示例 2:
输入: [3,4,-1,1]
输出: 2
示例 3:
输入: [7,8,9,11,12]
输出: 1
提示：
你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
*/

//这道常规写法不难，难在要求
func firstMissingPositive(nums []int) int {
	n := len(nums)
	//if n == 0 { // 这里可以不用判断
	//	return 1
	//}
	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}

	return n + 1
}
