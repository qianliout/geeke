package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNumArray_SumRange(t *testing.T) {
	convey.Convey("sum", t,
		func() {
			num := []int{1, 3, 5}
			n := Constructor(num)
			convey.So(n.SumRange(0, 2), convey.ShouldEqual, 9)
		})

	num := [][]int{{1, -1}, {2, 1}, {1, -2}, {2, 2}, {1, -3}, {2, 3}}
	sort.Slice(num, func(i, j int) bool {
		if num[i][0] != num[j][0] {
			return num[i][0] <= num[j][0]
		}
		return num[i][1] <= num[j][1]
	})
	fmt.Println(num)

}
