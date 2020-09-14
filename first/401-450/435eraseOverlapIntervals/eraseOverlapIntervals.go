package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	inter := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	ans := eraseOverlapIntervals(inter)
	fmt.Println("res is ", ans)
}

/*
给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
注意:
    可以认为区间的终点总是大于它的起点。
    区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。
示例 1:
输入: [ [1,2], [2,3], [3,4], [1,3] ]
输出: 1
解释: 移除 [1,3] 后，剩下的区间没有重叠。
示例 2:
输入: [ [1,2], [1,2], [1,2] ]
输出: 2
解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
*/

// 使用纯dp,先排序，这种方式是对的，但是是n2的复杂度，会超时-
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Sort(Item(intervals))
	fmt.Println(intervals)
	// dp[i] 表示先中第i个元素时的结果,有几个结果保留，
	dp := make(map[int]int)
	dp[0] = 1
	ans := 1

	for i := 0; i < len(intervals); i++ {
		for j := i - 1; j >= 0; j-- {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
				// break
			}
		}
		dp[i] = Max(dp[i], dp[i-1])
		ans = Max(ans, dp[i])

	}
	return len(intervals) - ans
}

type Item [][]int

func (ite Item) Len() int {
	return len(ite)
}

func (ite Item) Less(i, j int) bool {
	// 这里有个技巧，使用第二个数排序
	if ite[i][1] < ite[i][1] {
		return true
	} else if ite[i][1] == ite[i][1] {
		if ite[i][0] < ite[j][0] {
			return true
		}
	}
	return false
}

func (ite Item) Swap(i, j int) {
	ite[i], ite[j] = ite[j], ite[i]
}
func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
