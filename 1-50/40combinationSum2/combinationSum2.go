package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。
说明：
    所有数字（包括目标数）都是正整数。
    解集不能包含重复的组合。
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/
func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	path := make([]int, 0)
	sort.Ints(candidates)
	combinationHelper(candidates, path, target, 0, &res)
	return res
}

func combinationHelper(candidates, path []int, target, left int, res *[][]int) {
	if target == 0 {
		// fmt.Println("path is ", path)
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := left; i < len(candidates); i++ {
		if target < candidates[i] {
			return // 因为已排序的了
		}
		// 同层结点数值相同，剪枝.todo 这里一定要好好理解，把这里理解透了，dfs和回漱也理解透了
		if i != left && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		//下一次递归还是从i开始说明自己可以选
		combinationHelper(candidates, path, target-candidates[i], i+1, res)
		path = path[:len(path)-1]
		// path也可以在下一层再加，这样就不用回漱
		// combinationHelper(candidates, append(path, candidates[i]), target-candidates[i], i, res)
	}
}
