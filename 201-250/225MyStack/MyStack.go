package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.Push(1)
	m.Push(2)
	fmt.Println(m.Top())
	fmt.Println(m.Pop())
}

type MyStack struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	m := new(MyStack)
	m.input = make([]int, 0)
	m.output = make([]int, 0)
	return *m
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for _, v := range this.input {
		this.output = append(this.output, v)
	}
	this.input = make([]int, 0)
	if len(this.output) == 0 {
		return -1
	}
	v := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for _, v := range this.input {
		this.output = append(this.output, v)
	}
	this.input = make([]int, 0)
	if len(this.output) == 0 {
		return -1
	}
	return this.output[len(this.output)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	for _, v := range this.input {
		this.output = append(this.output, v)
	}
	this.input = make([]int, 0)
	return len(this.output) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
