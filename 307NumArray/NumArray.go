package main

import (
	"fmt"
)

/*
["NumArray","update","sumRange","sumRange","update","sumRange"]
[[[9,-8]],[0,3],[1,1],[0,1],[1,-3],[0,1]]
*/
func main() {
	num := []int{9, -8}
	arr := Constructor(num)
	arr.Update(0, 3)

	fmt.Println(arr.SumRange(1, 1))
	fmt.Println(arr.SumRange(0, 1))
	arr.Update(1, -3)
	fmt.Println(arr.SumRange(0, 1))
}

// 方法没有问题，但是会超出时间限制
type NumArray1 struct {
	Value []int
	Sum   int
	Left  []int
	Right []int
}

func Constructor1(nums []int) NumArray1 {
	left := make([]int, len(nums)+1)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		left[i+1] = left[i] + nums[i]
	}
	right := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = right[i+1] + nums[i]
	}

	return NumArray1{
		Value: nums,
		Sum:   sum,
		Left:  left,
		Right: right,
	}
}

func (this *NumArray1) SumRange(left int, right int) int {
	return this.Sum - this.Left[left] - this.Right[right+1]
}

func (this *NumArray1) Update(index int, val int) {
	if index < 0 || index >= len(this.Value) {
		return
	}
	sub := val - this.Value[index]
	for left := index + 1; left < len(this.Left); left++ {
		this.Left[left] += sub
	}
	for right := 0; right <= index; right++ {
		this.Right[right] += sub
	}

	this.Sum += sub

	this.Value[index] = val
}
