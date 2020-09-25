package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	res := minMoves2(nums)
	fmt.Println("res is ", res)
}

/*
给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。
您可以假设数组的长度最多为10000。
例如:
输入:
[1,2,3]
输出:
2
说明：
只有两个动作是必要的（记得每一步仅可使其中一个元素加1或减1）：
[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
*/
// 这题不用想什么中位数：设 a <= x <= b，将 a 和 b 都变化成 x 为最终目的，则需要步数为 x-a+b-x = b-a，即两个数最后相等的话步数一定是他们的差，x 在 a 和 b 间任意取；
// 所以最后剩的其实就是中位数；
// 那么直接排序后首尾指针计算就好：

func minMoves2(nums []int) int {
	sort.Ints(nums)
	ans := 0
	left, right := 0, len(nums)-1
	for left < right {
		ans += nums[right] - nums[left]
		left++
		right--
	}
	return ans
}
