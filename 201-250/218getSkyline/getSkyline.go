package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	. "outback/leetcode/common/heap/maxheap"
)

func main() {
	//buidings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	//buidings := [][]int{{0, 2, 3}, {2, 5, 3}}
	//buidings := [][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}}
	buidings := [][]int{{0, 3, 3}, {1, 5, 3}, {2, 4, 3}, {3, 7, 3}}
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

// 递归，分治和线段数，这里其实和线段数的解法没有多少关系
func getSkyline(buildings [][]int) [][]int {
	res := make([][]int, 0)
	if len(buildings) == 0 {
		return res
	}

	return segment(buildings, 0, len(buildings)-1)
}

func segment(building [][]int, l, r int) [][]int {
	res := make([][]int, 0)
	if l >= r {
		return [][]int{{building[l][0], building[l][2]}, {building[l][1], 0}}
	}
	mid := l + (r-l)/2
	left := segment(building, l, mid)
	right := segment(building, mid+1, r)

	// 合并，这里者是最难的
	m, n, lpre, rpre := 0, 0, 0, 0
	for m < len(left) || n < len(right) {
		if m >= len(left) {
			res = append(res, right[n])
			n++
		} else if n >= len(right) {
			res = append(res, left[m])
			m++
		} else {
			if left[m][0] == right[n][0] {
				if left[m][1] != Max(lpre, rpre) && left[m][1] >= right[n][1] {
					res = append(res, left[m])
				} else if left[m][1] <= right[n][1] && right[n][1] != Max(lpre, rpre) {
					res = append(res, right[n])
				}
				lpre = left[m][1]
				rpre = right[n][1]
				m++
				n++
			} else if left[m][0] < right[n][0] {
				if left[m][1] > rpre {
					res = append(res, left[m])
				} else if lpre > rpre {
					res = append(res, []int{left[m][0], rpre})
				}
				lpre = left[m][1]
				m++
			} else if left[m][0] > right[n][0] {
				if right[n][1] > lpre {
					res = append(res, right[n])
				} else if rpre > lpre {
					res = append(res, []int{right[n][0], lpre})
				}
				rpre = right[n][1]
				n++
			}
		}
	}
	return res
}

func Max(nums ...int) int {
	res := math.MinInt64
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res

}
