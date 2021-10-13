package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {

}

/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
示例：
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
说明:
    你可以假设数组不可变。
    会多次调用 sumRange 方法。
*/

type NumArray struct {
	Num    []int
	SumMap map[string]int
	format string
}

func Constructor(nums []int) NumArray {
	return NumArray{
		Num:    nums,
		SumMap: make(map[string]int),
		format: "%d_%d",
	}

}

func (this *NumArray) SumRange(i int, j int) int {
	if i >j {
		return 0
	}
	key := fmt.Sprintf(this.format, i, j)

	if v, ok := this.SumMap[key]; ok {
		return v
	}
	ans := 0
	minJ := common.Min(j, len(this.Num)-1)
	maxI := common.Max(i, 0)
	for k := maxI; k <= minJ; k++ {
		ans += this.Num[k]
	}

	this.SumMap[key] = ans
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
