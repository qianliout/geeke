package main

func main() {
}

/*
给定一个正整数数组 A，如果 A 的某个子数组中不同整数的个数恰好为 K，则称 A 的这个连续、不一定独立的子数组为好子数组。
（例如，[1,2,3,1,2] 中有 3 个不同的整数：1，2，以及 3。）
返回 A 中好子数组的数目。
示例 1：
输入：A = [1,2,1,2,3], K = 2
输出：7
解释：恰好由 2 个不同整数组成的子数组：[1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
示例 2：
输入：A = [1,2,1,3,4], K = 3
输出：3
解释：恰好由 3 个不同整数组成的子数组：[1,2,1,3], [2,1,3], [1,3,4].
*/

// 要计算（正好包含K个不同整数的子数组的）数量，可以计算（包含少于或等于K个不同整数的子数组的）数量，减去（包含少于或等于K-1个不同整数的子数组的）数量，来获得。而计算（包含少于或等于K个不同整数的子数组的）数量，即函数numberOfNotMoreThan(int K)，通过双指针实现，时间复杂度是O(n)。因此整体的时间复杂度是O(n)，与官方题解时间复杂度相同
func subarraysWithKDistinct(A []int, K int) int {
	ans := 0
	if len(A) < K {
		return ans
	}
	return numberOfNotMoreThan(A, K) - numberOfNotMoreThan(A, K-1)
}

// 少于或等k
func numberOfNotMoreThan(nums []int, k int) int {
	left, right, ans, numberOfDistinct := 0, 0, 0, 0
	need := make(map[int]int)
	for right < len(nums) {
		// 扩大窗口
		n := nums[right]
		need[n]++
		if need[n] == 1 {
			numberOfDistinct++
		}
		right++

		for numberOfDistinct > k && left < right {
			if need[nums[left]] == 1 {
				numberOfDistinct--
			}
			need[nums[left]]--
			left++
		}
		// 更新答案
		ans += right - left + 1
	}
	return ans
}

// 这样写是最简单的，但是，当数据量大一些时这里多一次遍历，会超时
func maplenght(m map[int]int) int {
	ans := 0
	for _, v := range m {
		if v > 0 {
			ans++
		}
	}
	return ans
}

func canAdd(m map[int]int, k int) bool {
	for _, v := range m {
		if v > 0 {
			k--
		}
	}
	if k == 0 {
		return true
	}
	return false
}
