package main

import (
	"fmt"

	. "outback/leetcode/common"
)

func main() {
	s := Constructor()
	s.AddNum(1)
	fmt.Println(s.GetIntervals())

	s.AddNum(3)
	fmt.Println(s.GetIntervals())
	s.AddNum(7)
	fmt.Println(s.GetIntervals())
	s.AddNum(2)
	fmt.Println(s.GetIntervals())

	s.AddNum(6)
	fmt.Println(s.GetIntervals())
}

/*
给定一个非负整数的数据流输入 a1，a2，…，an，…，将到目前为止看到的数字总结为不相交的区间列表。
例如，假设数据流中的整数为 1，3，7，2，6，…，每次的总结为：
[1, 1]
[1, 1], [3, 3]
[1, 1], [3, 3], [7, 7]
[1, 3], [7, 7]
[1, 3], [6, 7]
进阶：
如果有很多合并，并且与数据流的大小相比，不相交区间的数量很小，该怎么办?
*/

type SummaryRanges struct {
	bitsic *[]int
	used   *map[int]bool
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	b := make([]int, 0)
	used := make(map[int]bool)
	return SummaryRanges{
		bitsic: &b,
		used:   &used,
	}

}

func (this *SummaryRanges) AddNum(val int) {
	if !(*this.used)[val] {
		FindSmallIdxAndInsert(this.bitsic, val)
		(*this.used)[val] = true
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, 0)
	first, second := 0, 0
	for i := second + 1; i < len(*this.bitsic); i++ {
		if (*this.bitsic)[i] == (*this.bitsic)[second]+1 {
			second = i
		} else {
			res = append(res, []int{(*this.bitsic)[first], (*this.bitsic)[second]})
			first, second = second+1, second+1
		}
	}
	res = append(res, []int{(*this.bitsic)[first], (*this.bitsic)[second]})
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
