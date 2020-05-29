package main

func main() {
	nums := []int{1, 2, 3, 4}
	productExceptSelf(nums)
}

/*
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。
示例:
输入: [1,2,3,4]
输出: [24,12,8,6]
提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/
func productExceptSelf(nums []int) []int {
	// left,right表示left之前,或right之后的乘积,这里强调的是之前
	lenght := len(nums)
	res := make([]int, lenght)
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)

	left[0], left[1], right[lenght-2], right[lenght-1] = 1, nums[0], nums[lenght-1], 1

	for i := 1; i <= lenght; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := lenght - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < lenght; i++ {
		res[i] = left[i] * right[i]
	}

	//fmt.Print(left, right, res)
	return res
}
