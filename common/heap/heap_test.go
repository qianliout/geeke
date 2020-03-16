package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := make(IntMaxHeap, 0)
	maxHeap.Push(7)
	maxHeap.Push(3)
	maxHeap.Push(9)
	InitMax(&maxHeap)
	fmt.Println(maxHeap)

	//fmt.Println(maxHeap.PopMax())
	//fmt.Println(maxHeap.PopMax())
	//fmt.Println(maxHeap.PopMax())
}

func TestMinHeap(t *testing.T) {
	maxHeap := make(IntMinHeap, 0)
	maxHeap.Push(20)
	maxHeap.Push(22)
	maxHeap.Push(3)
	maxHeap.Push(9)
	InitMin(&maxHeap)
	fmt.Println(maxHeap)

	//fmt.Println(maxHeap.PopMax())
	//fmt.Println(maxHeap.PopMax())
	//fmt.Println(maxHeap.PopMax())
}
