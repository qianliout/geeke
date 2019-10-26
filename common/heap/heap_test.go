package heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := IntMaxHeap{}
	maxHeap.Push(4)
	maxHeap.Push(5)
	maxHeap.Push(6)
	InitMax(&maxHeap)
	fmt.Println(maxHeap.PopMax())
	fmt.Println(maxHeap.PopMax())
	fmt.Println(maxHeap.PopMax())
}
