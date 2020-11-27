package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	// nums := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	nums := [][]int{{1, 4}, {4, 5}}
	res := merge2(nums)
	fmt.Println(res)
}

/*
给出一个区间的集合，请合并所有重叠的区间。
示例 1:
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func merge(intervals [][]int) [][]int {
	sort.Sort(Item(intervals))

	res := make([][]int, 0)
	i := 0
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
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

func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(Item(intervals))
	res := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	i := 1
	for i < len(intervals) {

		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
		i++
		// 这一步是本道题的易错点
		if i == len(intervals) {
			res = append(res, []int{start, end})
		}
	}

	return res
}

type Item [][]int

func (item Item) Len() int {
	return len(item)
}

func (item Item) Less(i, j int) bool {
	if item[i][0] < item[j][0] {
		return true
	} else if item[i][0] == item[j][0] {
		if item[i][1] < item[j][1] {
			return true
		}
	}
	return false
}

func (item Item) Swap(i, j int) {
	item[i], item[j] = item[j], item[i]
}
