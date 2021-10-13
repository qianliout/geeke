package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	res := findMinArrowShots(points)
	fmt.Println("res is ", res)
}

/*
在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以y坐标并不重要，因此只要知道开始和结束的x坐标就足够了。开始坐标总是小于结束坐标。平面内最多存在104个气球。
一支弓箭可以沿着x轴从不同点完全垂直地射出。在坐标x处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
Example:
输入:
[[10,16], [2,8], [1,6], [7,12]]
输出:
2
解释:
对于该样例，我们可以在x = 6（射爆[2,8],[1,6]两个气球）和 x = 11（射爆另外两个气球）。
*/

// 这道题本质就是求有多少个不相交的区间
func findMinArrowShots(points [][]int) int {
	return useDP(points)
}

// 方法一，以起始点排序
func sortByFirst(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Sort(Item1(points))
	end := points[0][1]
	ans := 1

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][0]
			ans++
		} else {
			end = Min(end, points[i][1])
		}
	}
	return ans
}

// 方法二，以结束点排序
func sortByEnd(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	sort.Sort(item2(points))
	ans := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][1]
			ans++
		}
	}
	return ans
}

// 第二种方式，使用dp，但是时间复杂度是n2,不能ac
func useDP(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	sort.Sort(item2(points))
	dp := make(map[int]int)
	dp[0] = 1
	ans := 1
	for i := 1; i < len(points); i++ {
		for j := i - 1; j >= 0; j-- {
			if points[i][0] > points[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
				if dp[i] > ans {
					ans = dp[i]
				}
			}
		}
	}
	return ans
}

type Item1 [][]int

func (it Item1) Len() int {
	return len(it)
}

func (it Item1) Less(i, j int) bool {
	if it[i][0] > it[j][0] {
		return true
	}
	return false
}

func (it Item1) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

type item2 [][]int

func (it item2) Len() int {
	return len(it)
}

func (it item2) Less(i, j int) bool {
	if it[i][1] < it[j][1] {
		return true
	}
	return false
}

func (it item2) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
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
func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
