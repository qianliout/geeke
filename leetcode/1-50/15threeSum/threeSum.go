package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{0, 0, 0, 0, 0, 0}
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum3(nums)
	fmt.Println("ans is ", ans)
}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

	[[-1 0 1] 和 [0 1 -1]]是重复的三元组，
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	exit := make(map[string]struct{})
	backtrack(nums, path, 0, 0, &res, exit)
	return res
}

// 因为不可以包含重复的三元组,所以直接用回溯算法会出错
// 这种方法会超出时间限制
func backtrack(nums []int, path []int, start int, target int, res *[][]int, exit map[string]struct{}) {

	if len(path) > 3 {
		return
	}
	if target == 0 && len(path) == 3 {
		key := fmt.Sprintf("%d_%d_%d", path[0], path[1], path[2])
		if _, ok := exit[key]; !ok {
			ans := append([]int{}, path...)
			*res = append(*res, ans)
		}
		exit[key] = struct{}{}
		return
	}

	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && len(path) == 0 {
			continue
		}

		path = append(path, nums[i])
		target += nums[i]
		backtrack(nums, path, i+1, target, res, exit)
		path = path[:len(path)-1]
		target -= nums[i]
	}
}

func threeSum2(nums []int) [][]int {
	return Sum2(nums, 0)
}

func Sum2(nums []int, target int) [][]int {
	res := make([][]int, 0)

	if len(nums) <= 0 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s > target {
				right--
			} else if s < target {
				left++
			} else if s == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 这一步就是去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}

	for i := 0; i <= len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			tag := nums[i] + nums[left] + nums[right]
			if tag > 0 {
				right--
			} else if tag < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}
