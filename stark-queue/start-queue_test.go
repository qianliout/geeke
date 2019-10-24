package starkqueue

import (
	"fmt"
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	s := ")"
	r := IsValidParentheses(s)
	fmt.Println(r)
}
func TestMyQueue(t *testing.T) {
	myQueue := ConstructorQueue()
	myQueue.Push(5)
	myQueue.Push(4)
	n := myQueue.Pop()
	fmt.Println(n)
	v := myQueue.Peek()
	fmt.Println(v)
	fmt.Println(myQueue.Empty())
}
