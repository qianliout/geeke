package main

import (
	"sort"

	. "qianliout/leetcode/common/utils"
)

func main() {

}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回
一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))
	res := make([][]int, 0)
	if len(intervals) <= 1 {
		return intervals
	}
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			res = append(res, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = Max(right, intervals[i][1])
		}
	}
	res = append(res, []int{left, right})
	return res
}

type Intervals [][]int

func (it Intervals) Len() int {
	return len(it)
}

func (it Intervals) Less(i, j int) bool {
	if it[i][0] < it[j][0] {
		return true
	} else if it[i][0] > it[j][0] {
		return false
	} else if it[i][0] == it[j][0] {
		return it[i][1] < it[j][1]
	}
	return true
}

func (it Intervals) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
