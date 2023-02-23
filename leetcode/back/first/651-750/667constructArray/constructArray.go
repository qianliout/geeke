package main

func main() {

}

/*
给定两个整数 n 和 k，你需要实现一个数组，这个数组包含从 1 到 n 的 n 个不同整数，同时满足以下条件：
① 如果这个数组是 [a1, a2, a3, ... , an] ，那么数组 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数；.
② 如果存在多种答案，你只需实现并返回其中任意一种.
示例 1:
输入: n = 3, k = 1
输出: [1, 2, 3]
解释: [1, 2, 3] 包含 3 个范围在 1-3 的不同整数， 并且 [1, 1] 中有且仅有 1 个不同整数 : 1
示例 2:
输入: n = 3, k = 2
输出: [1, 3, 2]
解释: [1, 3, 2] 包含 3 个范围在 1-3 的不同整数， 并且 [2, 1] 中有且仅有 2 个不同整数: 1 和 2
提示:
 n 和 k 满足条件 1 <= k < n <= 104.
*/

func constructArray(n int, k int) []int {
	if n <= 0 || k <= 0 {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, 1)
	flag := true
	// 通过规律我们发现，要形成k个不同差值，那么这k个差值所形成的元素一定是驼峰型的
	// 即中间的数比两边的大，并且这个差值的绝对值一定是从k递减的
	// 并且相邻差值一定是一正一负
	for i := k; i > 1; i-- {
		if flag {
			res = append(res, res[len(res)-1]+i)
		} else {
			res = append(res, res[len(res)-1]-i)
		}
		flag = !flag
	}
	if k+2 <= n {
		for i := k + 2; i <= n; i++ {
			res = append(res, i)
		}
	}
	return res
}
