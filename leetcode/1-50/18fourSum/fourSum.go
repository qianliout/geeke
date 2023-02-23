package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	sum := fourSum(nums, 0)
	fmt.Println("fourSum is ", sum)
}

/*
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组
[nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案 。
示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]
提示：
1 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	if len(nums) < 4 {
		return ans
	}
	for i, b := range nums {
		if i > 0 && nums[i-1] == b {
			continue
		}
		sum3 := threeSum3(nums[i+1:], target-b)
		for j := range sum3 {
			ans = append(ans, append([]int{b}, sum3[j]...))
		}
	}
	return ans
}

func threeSum3(nums []int, target int) [][]int {
	// sort.Ints(nums)
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
			if tag > target {
				right--
			} else if tag < target {
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
