package heap

import (
	cheap "container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntMinHeap []int

func (h *IntMinHeap) Len() int           { return len(*h) }
func (h *IntMinHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntMinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 移出最小的那个值
func (h *IntMinHeap) PopMini() interface{} {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}

func (h *IntMinHeap) Peek() interface{} {
	x := (*h)[0]
	return x
}
func (h *IntMinHeap) PeekLast() interface{} {
	x := (*h)[len(*h)-1]
	return x
}

func InitMin(h *IntMinHeap) {
	cheap.Init(h)
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Example_intHeap() {
	h := &IntMinHeap{2, 1, 5}
	cheap.Init(h)
	cheap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", cheap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
