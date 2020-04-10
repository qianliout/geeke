package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/leetcode/common/heap/maxheap"
)

func main() {
	buidings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	//buidings := [][]int{{0, 2, 3}, {2, 5, 3}}
	//buidings := [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}
	//buidings := [][]int{{0, 3, 3}, {1, 5, 3}, {2, 4, 3}, {3, 7, 3}}
	res := getSkyline(buidings)
	fmt.Println("res is ", res)

}

// 方法一,大顶堆加 扫描法
func getSkyline2(buildings [][]int) [][]int {
	n := len(buildings)
	res := make([][]int, 0)
	pairs := make([][2]int, n*2)

	index := 0
	for _, build := range buildings {
		pairs[index][0] = build[0]
		pairs[index][1] = -build[2]
		index++
		pairs[index][0] = build[1]
		pairs[index][1] = build[2]
		index++
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] <= pairs[j][0]
		}
		return pairs[i][1] <= pairs[j][1]
	})

	maxHeap := make(MaxHeap, 0)
	prev := 0
	for _, pair := range pairs {
		if pair[1] < 0 {
			heap.Push(&maxHeap, -pair[1])
		} else {
			for i := 0; i < maxHeap.Len(); i++ {
				if maxHeap[i] == pair[1] {
					heap.Remove(&maxHeap, i)
					break
				}
			}
		}

		top := maxHeap.Peek().(int)
		if top != prev {
			res = append(res, []int{pair[0], top})
			prev = top
		}
	}
	return res
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	res := make([][]int, n)
	return res
}
