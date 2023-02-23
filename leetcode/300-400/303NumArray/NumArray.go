package main

func main() {

}

type NumArray struct {
	Value []int
	Sum   int
	Left  []int
	Right []int
}

func Constructor(nums []int) NumArray {
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

	return NumArray{
		Value: nums,
		Sum:   sum,
		Left:  left,
		Right: right,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Sum - this.Left[left] - this.Right[right+1]
}
