package main

import (
	"fmt"
)

func main() {
	fmt.Println(window([]int{10, 5, 2, 6}, 100))
	fmt.Println(window([]int{1, 1, 1, 1, 1}, 5))
}

/*
给定一个正整数数组 nums。
找出该数组内乘积小于 k 的连续的子数组的个数。
示例 1:
输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) <= 0 || k <= 0 {
		return 0
	}
	sum := 0
	// res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		// path2 := []int{nums[i]}
		// dfs(nums, 100, i, nums[i], path2, &sum, &res)
		dfs2(nums, k, i, nums[i], &sum)
	}
	// fmt.Println(res)
	return sum
}

// 注意的是连续的子数组
func dfs(nums []int, k, start, path int, path2 []int, target *int, res *[][]int) {
	if path < k {
		*target = *target + 1
		*res = append(*res, append([]int{}, path2...))
	} else {
		return
	}
	if start+1 < len(nums) {
		if path*nums[start+1] < k {
			path2 = append(path2, nums[start+1])
			dfs(nums, k, start+1, path*nums[start+1], path2, target, res)
			path2 = path2[:len(path2)-1]
		}
	}
}

// 这种方法可以用，但是会超时
func dfs2(nums []int, k, start, path int, target *int) {
	if path < k {
		*target = *target + 1
	} else {
		return
	}
	if start+1 < len(nums) && path*nums[start+1] < k {
		dfs2(nums, k, start+1, path*nums[start+1], target)
	}
}

// 使用滑动窗口的解法
func window(nums []int, k int) int {
	if len(nums) == 0 || k <= 1 {
		return 0
	}
	res, left, right, sum := 0, 0, 0, 1
	for right < len(nums) {
		sum *= nums[right]
		// 这里有没有left<=right这个条件都行，是为什么呢?,因为当left>right时，sum>=k这个条件一定不成立了
		for sum >= k && left <= right {
			sum /= nums[left]
			left++
		}
		res += right - left + 1
		right++
	}
	return res
}
