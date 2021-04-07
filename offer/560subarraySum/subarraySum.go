package main

func main() {

}

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :
数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
*/
func subarraySum(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	ans := 0
	for i := 1; i < len(preSum); i++ {
		for j := i - 1; j >= 0; j-- {
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}
	return ans
}

// 使用map
func subarraySum2(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1 // 空数组的前缀和为0,也可以这样去理解，当下面有个num等k时，就会去找之前等于0的前缀和，此时应该返回1才是正确答案
	ans, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans += preSum[sum-k]
		preSum[sum]++
	}
	return ans
}
