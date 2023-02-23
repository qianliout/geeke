package main

func main() {

}

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，
并且 i 和 j 的差的 绝对值 至多为 k。
示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}

	m := make(map[int]int)
	for i, v := range nums {
		if m[v] > 0 && i+1-m[v] <= k {
			return true
		}
		m[v] = i + 1 // 这里是易错点,因为map有默认值
	}
	return false
}
