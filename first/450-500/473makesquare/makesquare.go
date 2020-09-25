package main

func main() {

}

/*
还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。
输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。
示例 1:
输入: [1,1,2,2,2]
输出: true
解释: 能拼成一个边长为2的正方形，每边两根火柴。
示例 2:
输入: [3,3,3,3,4]
输出: false
解释: 不能用所有火柴拼成一个正方形。
注意:
    给定的火柴长度和在 0 到 10^9之间。
    火柴数组的长度不超过15。
*/

func makesquare(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 || len(nums) == 0 {
		return false
	}

	target := sum / 4
	used := make([]bool, len(nums))

	return backTrack(nums, &used, 4, target, 0, 0)

}

func backTrack(nums []int, selec *[]bool, k, target int, start, total int) bool {
	// 回溯算法 遍历整个决策树 寻找将nums拆分成4个子数组是否有可能
	// base case 当k为0时说明可以拼成4个和相同的子数组
	if k == 0 {
		return true
	}
	// 每次选择的数字总和等于target也就是正方形的边长时
	// 递归判断剩下的数组是否能拼成k-1个和为target的子数组
	// 这里start需要重置为0 但是有select数组在 所以不会重复选择元素

	if total == target {
		return backTrack(nums, selec, k-1, target, 0, 0)
	}

	for i := start; i < len(nums); i++ {
		if !(*selec)[i] && total+nums[i] <= target {
			(*selec)[i] = true
			if backTrack(nums, selec, k, target, i+1, total+nums[i]) {
				return true
			}
			(*selec)[i] = false
		}
	}
	return false
}
