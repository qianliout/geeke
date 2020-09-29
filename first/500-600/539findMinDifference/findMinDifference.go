package main

import (
	"math"
	"sort"
	"time"
)

func main() {

}

/*
给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
示例 1：
输入: ["23:59","00:00"]
输出: 1
备注:
列表中时间数在 2~20000 之间。
每个时间取值在 00:00~23:59 之间。
*/
func findMinDifference(timePoints []string) int {
	list := []int{}
	for _, s := range timePoints {
		parse, _ := time.Parse("2006-01-02T15:04", "2006-01-02T"+s)
		list = append(list, parse.Hour()*60+parse.Minute())
	}
	sort.Ints(list)
	ans := math.MaxInt64

	for i := 1; i < len(list); i++ {
		if list[i]-list[i-1] < ans {
			ans = list[i] - list[i-1]
		}
	}
	if 1440-list[len(list)-1]+list[0] < ans {
		ans = 1440 - list[len(list)-1] + list[0]
	}

	return ans
}
