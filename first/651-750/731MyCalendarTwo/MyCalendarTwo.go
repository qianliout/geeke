package main

func main() {

}

/*
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内不会导致三重预订时，则可以存储这个新的日程安排。
*/
type MyCalendar struct {
	timer Book // 这里不用排序也是可以的
}

func Constructor() MyCalendar {
	return MyCalendar{timer: make([][2]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	// 查看是否有相交
	for _, bk := range this.timer {
		if bk[0] >= end || bk[1] < start {
			continue
		} else {
			return false
		}
	}

	this.timer = append(this.timer, [2]int{start, end})
	return true
}

type Book [][2]int

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
