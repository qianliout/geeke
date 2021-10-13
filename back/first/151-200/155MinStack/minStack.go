package main

import (
	"container/heap"
	"math"

	"outback/leetcode/back/common/commonHeap"
)

func main() {

}

/*
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
示例:
输入：
["MinStackBySlice","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
输出：
[null,null,null,null,-3,null,0,-2]
解释：
MinStackBySlice minStack = new MinStackBySlice();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
提示：
pop、top 和 getMin 操作总是在 非空栈 上调用。
*/
type MinStackBySlice struct {
	minValue int
	value    []int
}

/** initialize your data structure here. */
func ConstructorBySlice() MinStackBySlice {
	return MinStackBySlice{minValue: math.MaxInt64, value: make([]int, 0)}
}

func (this *MinStackBySlice) Push(x int) {
	if x < this.minValue {
		this.minValue = x
	}
	this.value = append(this.value, x)
}

func (this *MinStackBySlice) Pop() {
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

func (this *MinStackBySlice) Top() int {
	if len(this.value) == 0 {
		return 0
	}
	return this.value[len(this.value)-1]
}

func (this *MinStackBySlice) GetMin() int {
	return this.minValue
}

/**
 * Your MinStackBySlice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// 使用最小堆来实现

type MinStack struct {
	minHeap commonHeap.MinHeap
	stack   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{minHeap: make(commonHeap.MinHeap, 0), stack: make([]int, 0)}
	return m
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	heap.Push(&this.minHeap, x)
}

func (this *MinStack) Pop() {
	n := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	for i := 0; i < len(this.minHeap); i++ {
		if this.minHeap[i] == n {
			heap.Remove(&this.minHeap, i)
			break
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	n := this.minHeap[0]
	return n
}
