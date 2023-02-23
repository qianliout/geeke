package main

import (
	"math"
)

func main() {

}

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
实现 MinStack 类:
MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。
*/

type MinStack struct {
	data    []int
	minData int
}

func Constructor() MinStack {
	return MinStack{data: make([]int, 0), minData: math.MaxInt64}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.minData = math.MaxInt64
	for i := range this.data {
		if this.data[i] < this.minData {
			this.minData = this.data[i]
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[:len(this.data)-1]
		this.minData = math.MaxInt64
		for i := range this.data {
			if this.data[i] < this.minData {
				this.minData = this.data[i]
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	return this.minData
}
