package commonHeap

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMaxHeap_Push(t *testing.T) {
	convey.Convey("max push", t, func() {
		maxh := make(MaxHeap, 0)
		heap.Push(&maxh, 9)
		heap.Push(&maxh, 12)
		heap.Push(&maxh, 3)
		heap.Push(&maxh, 29)
		heap.Push(&maxh, 3)
		min := maxh[0]
		fmt.Println(maxh)
		max := heap.Pop(&maxh)
		convey.So(min, convey.ShouldEqual, 29)
		convey.So(max, convey.ShouldEqual, 29)

	})

}
