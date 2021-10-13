package main

import (
	"outback/leetcode/back/common"
)

func main() {

}

/*
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
*/
type MyCalendarTwo struct {
	timer [][2]int
	over  [][2]int // 已经是两重预订的列表
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		timer: make([][2]int, 0),
		over:  make([][2]int, 0),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, bk := range this.over {
		if bk[0] >= end || bk[1] <= start {
			continue
		} else {
			return false
		}
	}
	this.addOver(start, end)
	this.timer = append(this.timer, [2]int{start, end})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func (this *MyCalendarTwo) addOver(start, end int) {
	for _, bk := range this.timer {
		if bk[0] >= end || bk[1] <= start {
			continue
		} else {
			this.over = append(this.over, [2]int{common.Max(start, bk[0]), common.Min(end, bk[1])})
		}
	}
}
