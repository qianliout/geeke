package starkqueue

/*
使用队列实现栈的下列操作：
    push(x) -- 元素 x 入栈
    pop() -- 移除栈顶元素
    top() -- 获取栈顶元素
    empty() -- 返回栈是否为空
注意:
    你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
    你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
    你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
*/

type InQueue []int
type OutQueue []int
type MyStack struct {
	Input  InQueue
	Output OutQueue
}

func ConstructorStark() *MyStack {
	return &MyStack{
		Input:  make(InQueue, 0),
		Output: make(OutQueue, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Input = append(this.Input, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for _, value := range this.Input {
		this.Output = append(this.Output, value)
	}
	this.Input = make(InQueue, 0)
	if len(this.Output) > 0 {
		value := this.Output[0]
		this.Output = this.Output[1:]
		return value
	}
	return 0
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for _, value := range this.Input {
		this.Output = append(this.Output, value)
	}
	this.Input = make(InQueue, 0)
	if len(this.Output) > 0 {
		value := this.Output[0]
		return value
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	for _, value := range this.Input {
		this.Output = append(this.Output, value)
	}
	this.Input = make(InQueue, 0)
	if len(this.Output) > 0 {
		return false
	}
	return true
}
