package main

func main() {

}

/*
给定一个未排序的整数数组，找到最长递增子序列的个数。
输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:
输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。
*/
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// dp[i]表示以ii位置结尾，即nums[i]nums[i]值结尾的，最长连续递增序列的长度
	dp := make([]int, len(nums))
	// counter[i]表示以ii位置结尾的最长连续递增序列的个数
	counter := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		counter[i] = 1
	}
	max := 0
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					counter[i] = counter[j]
				} else if dp[j]+1 == dp[i] {
					counter[i] += counter[j]
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == max {
			ans += counter[i]
		}
	}
	return ans
}
