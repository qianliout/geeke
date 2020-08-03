package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	res := fourSum(nums, 0)
	fmt.Println(res)
}

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	exitMap := make(map[int]int)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				s := nums[i] + nums[j] + nums[k]
				c := target - s
				if n, ok := exitMap[c]; ok {
					res = append(res, []int{nums[i], nums[j], nums[k], n})
				}
				exitMap[s] = s
			}
		}
	}
	return res
}
