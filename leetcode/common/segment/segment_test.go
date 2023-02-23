package segment

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewSegment(t *testing.T) {
	convey.Convey("new segent", t,
		func() {
			nums := []int{2, 4, 5, 7, 8, 9}
			s := NewSegment(nums)
			tree := []int{0, 35, 29, 6, 12, 17, 2, 4, 5, 7, 8, 9}
			convey.So(s.Data, convey.ShouldEqual, nums)
			convey.So(s.Tree, convey.ShouldEqual, tree)
		})
}

func TestSegment_Update(t *testing.T) {

	convey.Convey("update", t,
		func() {
			nums := []int{2, 4, 5, 7, 8, 9}
			s := NewSegment(nums)
			s.Update(0, 3)
			tree := []int{0, 35, 29, 6, 12, 17, 2, 4, 5, 7, 8, 9}
			// convey.So(s.Data, convey.ShouldEqual, nums)
			convey.So(s.Tree, convey.ShouldEqual, tree)
		})

}

func TestSegment_SumRange(t *testing.T) {

	convey.Convey("sum ", t,
		func() {
			nums := []int{1, 3, 5}
			s := NewSegment(nums)
			sum := s.SumRange(0, 2)
			fmt.Println(s.Tree)
			convey.So(sum, convey.ShouldEqual, 9)

		})
}
