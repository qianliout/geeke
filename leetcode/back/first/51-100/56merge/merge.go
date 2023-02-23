package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]int{{1, 4}, {1, 4}, {8, 10}, {15, 18}}
	res := merge2(in)
	fmt.Println(res)
}

// 一种方法是罗列出所有的情况，这种方法不难，只是要小心，把各种情况列全
func merge(intervals [][]int) [][]int {
	start := 1
	for start < len(intervals) {
		/*
			第一种情况：【1 3】，【2，4】
			第二种情况：【1 4】，【2，3】
			第三种情况：【2 3】，【1，4】
			第四种情况 【1 4】，【0 1】
			第五种情况：【1 2】，【3，4】，这种情况才把start+1
		*/
		if intervals[start-1][0] <= intervals[start][0] && intervals[start-1][1] <= intervals[start][1] && intervals[start-1][1] >= intervals[start][0] {
			intervals[start-1][1] = intervals[start][1]                   // 先合并i-1个节点
			intervals = append(intervals[:start], intervals[start+1:]...) // 删除第i个节点
		} else if intervals[start-1][0] >= intervals[start][0] && intervals[start-1][1] >= intervals[start][1] && intervals[start-1][0] >= intervals[start][1] {
			intervals[start-1][0] = intervals[start][0]                   // 先合并i-1个节点
			intervals = append(intervals[:start], intervals[start+1:]...) // 删除第i个节点
		} else if intervals[start-1][0] <= intervals[start][0] && intervals[start-1][1] >= intervals[start][1] {
			intervals = append(intervals[:start], intervals[start+1:]...) // 删除第i个节点
		} else if intervals[start-1][0] >= intervals[start][0] && intervals[start-1][1] <= intervals[start][1] {
			intervals = append(intervals[:start-1], intervals[start:]...) // 删除第i个节点
		} else {
			start++
		}
	}
	return intervals
}

func merge2(intervals [][]int) [][]int {
	// 先排序
	sort.Sort(Tupe(intervals))
	start := 1

	for start < len(intervals) {
		if intervals[start-1][0] <= intervals[start][0] && intervals[start-1][1] >= intervals[start][0] {
			if intervals[start-1][1] <= intervals[start][1] {
				intervals[start-1][1] = intervals[start][1]                   // 先合并i-1个节点
				intervals = append(intervals[:start], intervals[start+1:]...) // 删除第i个节点
			} else {
				intervals = append(intervals[:start], intervals[start+1:]...) // 删除第i个节点
			}
		} else {
			start++
		}
	}
	return intervals
}

type Tupe [][]int

func (p Tupe) Len() int { return len(p) }
func (p Tupe) Less(i, j int) bool {
	if p[i][0] < p[j][0] {
		return true
	}
	if p[i][0] == p[j][0] && p[i][1] < p[j][1] {
		return true
	}
	return false
}

func (p Tupe) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
