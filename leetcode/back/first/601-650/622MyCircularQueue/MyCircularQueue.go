package main

import (
	"fmt"

	listnode2 "qianliout/leetcode/leetcode/common/listnode"
)

func main() {
	// ["MyCircularQueue","enQueue","Rear","Rear","deQueue","enQueue","Rear","deQueue","Front","deQueue","deQueue","deQueue"]
	// [[6],[6],[],[],[],[5],[],[],[],[],[],[]]
	// ["MyCircularQueue","enQueue","deQueue","Front","deQueue","Front","Rear","enQueue","isFull","deQueue","Rear","enQueue"]
	// [[3],[7],[],[],[],[],[],[0],[],[],[],[3]]
	obj := Constructor(3)
	fmt.Println(obj.EnQueue(7))
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())

	fmt.Println(obj.Rear())
	fmt.Println(obj.EnQueue(0))
	fmt.Println(obj.IsFull())

	fmt.Println(obj.DeQueue())
	// fmt.Println(obj.Front())
	fmt.Println(obj.Rear())
	fmt.Println(obj.EnQueue(3))
}

type MyCircularQueue struct {
	length   int
	capacity int
	head     *listnode2.ListNode
	tail     *listnode2.ListNode
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	c := new(MyCircularQueue)
	c.capacity = k
	return *c
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.length >= this.capacity {
		return false
	}
	n := listnode2.ListNode{Val: value}
	if this.length == 0 {
		this.head = &n
		this.tail = this.head
	} else {
		this.tail.Next = &n
		this.tail = &n
	}
	this.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.length <= 0 {
		return false
	}
	this.head = this.head.Next
	this.length--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.length == 0 {
		return -1
	}
	return this.head.Val
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.length == 0 {
		return -1
	}
	return this.tail.Val
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
