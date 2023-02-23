package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(candidates, 8)
	fmt.Println("res is ", res)

}

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
*/

// 本题目的难点是，数组有重复的元素，但是不能有重复解
func combinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	path := make([]int, 0)
	res := make([][]int, 0)
	backtracking(candidates, path, target, 0, &res)
	return res
}

func backtracking(candidates []int, path []int, target int, start int, res *[][]int) {

	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		// 同层结点数值相同，剪枝
		if i != start && candidates[i] == candidates[i-1] {
			continue
		}

		target -= candidates[i]
		path = append(path, candidates[i])
		// 如果数组中的每一个元素只能使用多次，就是i
		backtracking(candidates, path, target, i+1, res)
		target += candidates[i]
		path = path[:len(path)-1]
	}
}
