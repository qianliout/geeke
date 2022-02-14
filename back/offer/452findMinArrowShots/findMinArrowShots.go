package main

import (
	"sort"

	"qianliout/leetcode/common/utils"
)

func main() {

}

/*
在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。
一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
*/

// 不就是找不连续的区间吗
func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	ans := 1
	sort.Sort(Item(points))
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			ans++
			end = points[i][1]
		} else {
			end = utils.Min(end, points[i][1])
		}
	}
	return ans
}

// 会超时
func useDP(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	dp := make(map[int]int)
	dp[0] = 1
	ans := 1
	sort.Sort(Item(points))
	for i := 1; i < len(points); i++ {
		for j := i - 1; j >= 0; j-- {
			if points[i][0] > points[j][1] {
				dp[i] = utils.Max(dp[i], dp[j]+1)
			}
		}
		dp[i] = utils.Max(dp[i], dp[i-1])
		ans = utils.Max(dp[i], ans)
	}
	return ans
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
