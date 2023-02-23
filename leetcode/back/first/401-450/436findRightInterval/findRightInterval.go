package main

import (
	"fmt"
	"sort"
)

func main() {
	inter := [][]int{{1, 2}, {2, 3}, {0, 1}, {3, 4}}
	res := findRightInterval(inter)
	fmt.Println("res is ", res)
}

/*
给定一组区间，对于每一个区间 i，检查是否存在一个区间 j，它的起始点大于或等于区间 i 的终点，这可以称为 j 在 i 的“右侧”。
对于任何区间，你需要存储的满足条件的区间 j 的最小索引，这意味着区间 j 有最小的起始点可以使其成为“右侧”区间。如果区间 j 不存在，
则将区间 i 存储为 -1。最后，你需要输出一个值为存储的区间值的数组。
注意:
你可以假设区间的终点总是大于它的起始点。
你可以假定这些区间都不具有相同的起始点。
示例 1:
输入: [ [1,2] ]
输出: [-1]
解释:集合中只有一个区间，所以输出-1。
示例 2:
输入: [ [3,4], [2,3], [1,2] ]
输出: [-1, 0, 1]
解释:对于[3,4]，没有满足条件的“右侧”区间。
对于[2,3]，区间[3,4]具有最小的“右”起点;
对于[1,2]，区间[2,3]具有最小的“右”起点。
示例 3:
输入: [ [1,4], [2,3], [3,4] ]
输出: [-1, 2, -1]
解释:对于区间[1,4]和[3,4]，没有满足条件的“右侧”区间。
对于[2,3]，区间[3,4]有最小的“右”起点。
*/

func findRightInterval(intervals [][]int) []int {
	// 记下起始点的索引
	ans := make([]int, len(intervals))
	if len(intervals) == 0 {
		return ans
	}
	if len(intervals) == 0 {
		ans[0] = -1
		return ans
	}

	indexMap := make(map[int]int)
	for i, v := range intervals {
		indexMap[v[0]] = i
	}
	sort.Sort(MaxHeap(intervals))
	i := 0
	for i < len(intervals)-1 {
		ans[indexMap[intervals[i][0]]] = -1
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][1] <= intervals[j][0] {
				ans[indexMap[intervals[i][0]]] = indexMap[intervals[j][0]]
				break
			}
		}
		i++
	}
	ans[indexMap[intervals[i][0]]] = -1
	return ans
}

type MaxHeap [][]int

func (it MaxHeap) Len() int {
	return len(it)
}

func (it MaxHeap) Less(i, j int) bool {
	// 你可以假定这些区间都不具有相同的起始点 所以这样就能判断了
	return it[i][0] < it[j][0]
}

func (it MaxHeap) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Peek() []int {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return []int{}
}
