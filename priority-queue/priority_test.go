package priority

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	//kth := ConstructorKthLargest(3, []int{3, 4, 5, 6, 7})
	kth := ConstructorKthLargestByHeap(3, []int{3, 4, 5, 6, 7})
	fmt.Println(kth.Add(8))
}
