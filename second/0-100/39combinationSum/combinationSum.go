package main

import (
	"fmt"
)

func main() {
	num := []int{3, 2, 5}
	res := combinationSum(num, 8)
	fmt.Println("res is ", res)
}

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5,3], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
提示：
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	dfs(candidates, []int{}, target, 0, &res)
	return res
}

func dfs(candidates, path []int, left, start int, res *[][]int) {
	if left < 0 {
		return
	}
	if left == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	if left > 0 {
		for i := start; i < len(candidates); i++ {
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			dfs(candidates, append(path, candidates[i]), left-candidates[i], i, res)
		}
	}
}
