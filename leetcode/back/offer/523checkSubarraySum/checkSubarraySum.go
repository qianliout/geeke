package main

func main() {

}

/*
给定一个包含 非负数 的数组和一个目标 整数 k ，编写一个函数来判断该数组是否含有连续的子数组，
其大小至少为 2，且总和为 k 的倍数，即总和为 n * k ，其中 n 也是一个整数。
示例 1：
输入：[23,2,4,6,7], k = 6
输出：True
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6。
示例 2：
输入：[23,2,6,4,7], k = 6
输出：True
解释：[23,2,6,4,7]是大小为 5 的子数组，并且和为 42。
*/
// 两次循环
func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i-1; j++ { // 因为至少有两个，所以是i-1
			if (k != 0 && (preSum[i]-preSum[j])%k == 0) || (k == 0 && preSum[i]-preSum[j] == 0) {
				return true
			}
		}
	}
	return false
}

// 前缀和
// sum[i:j] = sum[:j] - sum[:i]
// 可以被k整除的连续数组，只要保证sum[:j]和sum[:i]的对k的余数相等即可。
func checkSubarraySum2(nums []int, k int) bool {
	preSum := make(map[int]int)
	preSum[0] = -1 // 这个怎么理解
	pre := 0
	for i, v := range nums {
		pre += v
		// 这一步也容易出错，就是怕k是0的情况
		if k != 0 {
			pre = pre % k
		}

		preInx, ok := preSum[pre]
		if ok && i-preInx >= 2 {
			return true
		}
		if !ok {
			// 这里一定要注意，只是不存在时才更新，就是怕有连续0的情况，也是为了保证最长
			preSum[pre] = i
		}
	}
	return false
}
