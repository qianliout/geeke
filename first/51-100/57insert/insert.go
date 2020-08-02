package main

import (
	"fmt"
	"math"
)

/*
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
示例 1:
输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
输出: [[1,5],[6,9]]
示例 2:
输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出: [[1,2],[3,10],[12,16]]
解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
*/

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	res := insert(intervals, newInterval)
	fmt.Println(res)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	output := make([][]int, 0)

	newSart, newEnd := newInterval[0], newInterval[1]
	idx, n := 0, len(intervals)
	for idx < n && newSart > intervals[idx][0] {
		output = append(output, intervals[idx])
		idx++
	}
	if len(output) <= 0 || output[len(output)-1][1] < newSart {
		output = append(output, newInterval)
	} else {
		output[len(output)-1][1] = int(math.Max(float64(output[len(output)-1][1]), float64(newEnd)))
	}
	// 继序合并
	for idx < n {
		interval := intervals[idx]
		start, end := interval[0], interval[1]

		if output[len(output)-1][1] < start {
			output = append(output, interval)
		} else {
			output[len(output)-1][1] = int(math.Max(float64(output[len(output)-1][1]), float64(end)))
		}
		idx++
	}
	return output
}
