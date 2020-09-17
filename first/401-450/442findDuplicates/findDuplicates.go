package main

func main() {

}

/*
给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
找到所有出现两次的元素。
你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？
示例：
输入:
[4,3,2,7,8,2,3,1]
输出:
[2,3]
*/

func findDuplicates(nums []int) []int {
	return userMap(nums)
}

// 使用map

func userMap(nums []int) []int {
	res := make([]int, 0)
	exit := make(map[int]int)
	for _, n := range nums {
		exit[n]++
	}
	for k, v := range exit {
		if v == 2 {
			res = append(res, k)
		}
	}
	return res
}

// 取反做标记，因为我题目中说了1 ≤ a[i] ≤ n,所以可以直接取反

func negativeMarker(nums []int) []int {
	res := make([]int, 0)
	for _, n := range nums {
		index := n - 1
		if n <= 0 {
			index = -n - 1
		}

		if nums[index] > 0 {
			nums[index] = -nums[index]
		} else {
			res = append(res, index+1)
		}
	}
	return res
}
