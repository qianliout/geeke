package main

func main() {

}

/*
给定一个迭代器类的接口，接口包含两个方法： next() 和 hasNext()。设计并实现一个支持 peek()
操作的顶端迭代器 -- 其本质就是把原本应由 next() 方法返回的元素 peek() 出来。
示例:
假设迭代器被初始化为列表 [1,2,3]。
调用 next() 返回 1，得到列表中的第一个元素。
现在调用 peek() 返回 2，下一个元素。在此之后调用 next() 仍然返回 2。
最后一次调用 next() 返回 3，末尾元素。在此之后调用 hasNext() 应该返回 false。
链接：https://leetcode-cn.com/problems/peeking-iterator
*/

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	Queue []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	queue := make([]int, 0)
	for iter.hasNext() {
		queue = append(queue, iter.next())
	}
	return &PeekingIterator{Queue: queue}
}

func (this *PeekingIterator) hasNext() bool {
	if len(this.Queue) > 0 {
		return true
	}
	return false

}

func (this *PeekingIterator) next() int {
	n := this.Queue[0]
	this.Queue = this.Queue[1:len(this.Queue)]
	return n
}

func (this *PeekingIterator) peek() int {
	return this.Queue[0]
}
