package main

import (
	"fmt"
	"sort"
)

/*
  给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
  candidates 中的数字可以无限制重复被选取。
  说明：
    所有数字（包括 target）都是正整数。
    解集不能包含重复的组合。
  示例 1:
  输入: candidates = [2,3,6,7], target = 7,
  所求解集为:
  [
  [7],
  [2,2,3]
  ]
  示例 2:
  输入: candidates = [2,3,5], target = 8,
  所求解集为:
  [
  [2,2,2,2],
  [2,3,3],
  [3,5]
  ]
*/

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	if len(candidates) == 0 {
		return res
	}
	path := make([]int, 0)
	sort.Ints(candidates) // 通过提交可以知道，排序可以加快速度
	combinationHelper(candidates, path, target, 0, &res)
	return res
}

func combinationHelper(candidates, path []int, target, left int, res *[][]int) {
	if target == 0 {
		fmt.Println("path is ", path)
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := left; i < len(candidates); i++ {
		if target < candidates[i] {
			return // 因为已排序的了
		}
		path = append(path, candidates[i])
		combinationHelper(candidates, path, target-candidates[i], i, res)
		path = path[:len(path)-1]
		// path也可以在下一层再加，这样就不用回漱
		// combinationHelper(candidates, append(path, candidates[i]), target-candidates[i], i, res)
	}
}

func sumSlice(res []int) int {
	s := 0
	for _, v := range res {
		s += v
	}
	return s
}
