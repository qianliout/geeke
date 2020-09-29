package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0}
	res := check(nums, 0)

	fmt.Println("res is ", res)
}

/*
给定一个包含 非负数 的数组和一个目标 整数 k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，且总和为 k 的倍数，即总和为 n*k，其中 n 也是一个整数。
示例 1：
输入：[23,2,4,6,7], k = 6
输出：True
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6。
示例 2：
输入：[23,2,6,4,7], k = 6
输出：True
解释：[23,2,6,4,7]是大小为 5 的子数组，并且和为 42。
说明：
数组的长度不会超过 10,000 。
你可以认为所有数字总和在 32 位有符号整数范围内。
*/

// [0,0] 0 这种结果也得注意

// 我们将nums的数据拷贝到dp中
// 计算长度为2的子数组和时：dp[j] = dp[j] + nums[j+1]; 这里的dp[j]就是nums[j]
// 计算长度为3的子数组和时：dp[j] = dp[j] + nums[j+2]; 这里的dp[j]是更新过的dp[j]，一个dp[j]相当于nums[j]+nums[j+1]
// 这样当计算长度为p的子数组大小时，就可以利用已经计算过的长度为p-1的子数组进行更新，就可以对原来的三重循环进行优化，变为二重循环
func checkSubarraySum(nums []int, k int) bool {
	// dp表法到的元素和，包括i,不包括j
	if len(nums) <= 1 {
		return false
	}
	// dp[j]表示，以j开始，2，3，4...的数之和的余数
	dp := make(map[int]int)

	if k == 0 {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-i; j++ {
				dp[j] = dp[j] + nums[j+i]
				if i != 0 && dp[j] == 0 {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			dp[j] = (dp[j] + nums[j+i]) % k
			if i != 0 && dp[j] == 0 {
				return true
			}
		}
	}
	return false
}

// 前缀和
// sum[i:j] = sum[:j] - sum[:i]
// 可以被k整除的连续数组，只要保证sum[:j]和sum[:i]的对k的余数相等即可。
func check(nums []int, k int) bool {
	dp := make(map[int]int) // key是求前缀和的余数，int是index,也就是nums[:i]
	pre := 0
	dp[0] = -1 // 主要是为了应付:[0,0],0这种情况
	for i, n := range nums {
		pre += n
		if k != 0 {
			pre = pre % k
		}
		idx, ok := dp[pre]
		if ok {
			if i-idx > 1 {
				return true
			}
		} else {
			dp[pre] = i
		}
	}
	return false
}
