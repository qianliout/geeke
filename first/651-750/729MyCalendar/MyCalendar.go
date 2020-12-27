package main

func main() {

}

/*
实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false 并且不要将该日程安排添加到日历中。
请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar(); MyCalendar.book(start, end)
示例 1:
MyCalendar();
MyCalendar.book(10, 20); // returns true
MyCalendar.book(15, 25); // returns false
MyCalendar.book(20, 30); // returns true
解释:
第一个日程安排可以添加到日历中.  第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了。
第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 。
说明:
每个测试用例，调用 MyCalendar.book 函数最多不超过 100次。
调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
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

func (bk Book) Len() int {
	return len(bk)
}

func (bk Book) Less(i, j int) bool {
	if bk[i][0] < bk[j][0] {
		return true
	} else if bk[i][0] == bk[j][0] {
		return bk[i][1] < bk[j][1]
	}
	return false
}

func (bk Book) Swap(i, j int) {
	bk[i], bk[j] = bk[j], bk[i]
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
