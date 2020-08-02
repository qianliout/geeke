package main

import (
	"fmt"

	. "outback/leetcode/common/segment"
)

func main() {

}

/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
update(i, val) 函数可以通过将下标为 i 的数值更新为 val，从而对数列进行修改。
示例:
Given nums = [1, 3, 5]
sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
说明:
    数组仅可以在 update 函数下进行修改。
    你可以假设 update 函数与 sumRange 函数的调用次数是均匀分布的。
*/

type NumArray struct {
	Segment Segment
}

func Constructor(nums []int) NumArray {
	n := new(NumArray)
	n.Segment = NewSegment(nums)
	return *n
}

func (this *NumArray) Update(i int, val int) {
	this.Segment.Update(i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	fmt.Println(this.Segment.Tree)
	return this.Segment.SumRange(i, j)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
