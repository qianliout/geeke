package priority

import (
	"fmt"
	"testing"

	"outback/leetcode/common/commonHeap"
)

func TestKthLargest(t *testing.T) {
	//kth := ConstructorKthLargest(3, []int{3, 4, 5, 6, 7})
	kth := ConstructorKthLargestByHeap(3, []int{3, 4, 5, 6, 7})
	fmt.Println(kth.Add(5))
}

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := MaxSlidingWindow(nums, 3)
	fmt.Println(res)
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
