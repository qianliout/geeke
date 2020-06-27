package main

import (
	"math"
)

func main() {

}

type MinStack struct {
	minValue int
	value    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{minValue: math.MaxInt64, value: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if x < this.minValue {
		this.minValue = x
	}
	this.value = append(this.value, x)
}

func (this *MinStack) Pop() {
	if len(this.value) == 0 {
		return
	}
	if this.value[len(this.value)-1] == this.minValue {
		min := math.MaxInt64
		for i := 0; i < len(this.value)-1; i++ {
			if this.value[i] < min {
				min = this.value[i]
			}
		}
		this.minValue = min
	}
	this.value = this.value[:len(this.value)-1]
}

func (this *MinStack) Top() int {
	if len(this.value) == 0 {
		return 0
	}
	return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
