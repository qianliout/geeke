package priority

import (
	"fmt"
	commonHeap "outback/leetcode/common/heap"
	"testing"
)

func TestKthLargest(t *testing.T) {
	//kth := ConstructorKthLargest(3, []int{3, 4, 5, 6, 7})
	kth := ConstructorKthLargestByHeap(3, []int{3, 4, 5, 6, 7})
	fmt.Println(kth.Add(1))
}

func TestMinHeap(t *testing.T) {
	min := commonHeap.IntMinHeap{}
	min.Push(3)
	min.Push(4)
	min.Push(5)
	fmt.Println(min.Peek())
	fmt.Println(min)
	fmt.Println(min.Pop())
	fmt.Println(min)
}
