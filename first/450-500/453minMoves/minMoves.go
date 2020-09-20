package main

import (
	"math"
)

func main() {

}

/*
给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动将会使 n - 1 个元素增加 1。
示例:
输入:
[1,2,3]
输出:
3
解释:
只需要3次移动（注意每次移动会增加两个元素的值）：
[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
*/

// 每次使n-1个元素+1 完全等效于 每次使1个元素-1
// 所以只需找出：每次让1个元素-1，多少次后所有元素相等
// 这不就简单了吗？找出最小值然后遍历一遍，over！
func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := math.MaxInt64

	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	ans := 0
	for _, n := range nums {
		ans += n - min
	}
	return ans
}
