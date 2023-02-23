package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{10, 4}, {13, 18}, {1, 5}, {13, 15}, {3, 12}, {12, 11}, {17, 15}, {7, 1}, {17, 18}, {7, 19}, {2, 5}, {8, 9}, {18, 10}, {7, 6}, {17, 7}}
	res := maxEnvelopes(nums)

	fmt.Println(res)
}

/*
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，
这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
说明:
不允许旋转信封。
示例:
输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
*/

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 || len(envelopes[0]) == 0 {
		return 0
	}
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Sort(Item(envelopes))
	fmt.Println(envelopes)
	// dp[i]表示i这个信峰在最外面时一个可以套多少个了
	dp := make(map[int]int)
	dp[0] = 1
	max := dp[0]
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

type Item [][]int

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	if it[i][0] < it[j][0] {
		return true
	} else if it[i][0] > it[j][0] {
		return false
	} else {
		if it[i][1] < it[j][1] {
			return true
		} else {
			return false
		}
	}
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
