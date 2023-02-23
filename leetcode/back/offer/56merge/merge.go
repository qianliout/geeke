package main

import (
	"sort"
)

func main() {

}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Sort(Item(intervals))
	left := 1
	if len(intervals) <= 1 {
		return intervals
	}
	start := intervals[0][0]
	end := intervals[0][1]
	for left < len(intervals) {
		if intervals[left][0] <= end {
			if intervals[left][1] > end {
				end = intervals[left][1]
			}
		} else {
			ans = append(ans, []int{start, end})
			start = intervals[left][0]
			end = intervals[left][1]
		}
		left++
	}
	ans = append(ans, []int{start, end})
	return ans
}

func merge2(intervals [][]int) [][]int {
	sort.Sort(Item(intervals))

	res := make([][]int, 0)
	i := 0
	for i < len(intervals) {
		start, end := intervals[i][0], intervals[i][1]
		i++
		for i < len(intervals) && end >= intervals[i][0] {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
			i++
		}
		res = append(res, []int{start, end})
	}
	return res
}

type Item [][]int

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if it[i][0] < it[j][0] || (it[i][0] == it[j][0] && it[i][1] < it[i][1]) {
		return true
	}
	return false
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
