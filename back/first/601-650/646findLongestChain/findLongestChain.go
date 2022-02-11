package main

import (
	"sort"

	common2 "qianliout/leetcode/common"
)

func main() {

}

/*
给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
示例：
输入：[[1,2], [2,3], [3,4]]
输出：2
解释：最长的数对链是 [1,2] -> [3,4]
*/

func findLongestChain(pairs [][]int) int {
	// 这一步判断可以不要
	if len(pairs) <= 1 {
		return 0
	}
	sort.Sort(item(pairs))
	start := pairs[0][1]
	ans := 1
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > start {
			ans++
			start = pairs[i][1]
		} else {
			start = common2.Min(start, pairs[i][1])
		}
	}
	return ans
}

type item [][]int

func (it item) Len() int {
	return len(it)
}

func (it item) Less(i, j int) bool {
	if it[i][0] < it[j][0] {
		return true
	} else if it[i][0] == it[j][0] && it[i][1] < it[j][1] {
		return true
	}
	return false
}

func (it item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
