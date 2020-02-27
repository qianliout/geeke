package main

//func permute5(nums []int) [][]int {
//	res = make([][]int, 0)
//	var backtrack func(nums, tep []int)
//
//	backtrack = func(ints, tep []int) {
//		if len(ints) == 0 {
//			res = append(res, tep)
//			return
//		}
//		for i := 0; i < len(ints); i++ {
//			tep = append(tep, ints[i])
//			ints := append(ints[:i], ints[i+1:]...)
//			backtrack(ints, tep)
//		}
//	}
//	tem := make([]int, 0)
//	backtrack(nums, tem)
//	return res
//}
//var res [][]int

func permute1(nums []int) [][]int {
	size := len(nums)
	if size == 0 {
		return nil
	}
	ans := make([]int, 0)
	res := make([][]int, 0)
	used := make([]bool, size)
	helper(nums, used, ans, &res)

	return res
}

func helper(nums []int, used []bool, ans []int, res *[][]int) {
	if len(ans) == len(nums) {
		*res = append(*res, append([]int{}, ans...))
		//fmt.Println("ans ", ans)
		/*
		   注意这里很容易出错，通过打印发现，每次执行到这里时ans是正确的值，但是res append是ans的内存地址，
		   但是递归回来时会把这个内存地址的值改了，所以一定要注意
		*/
		//*res = append(*res, ans)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			ans = append(ans, nums[i])
			helper(nums, used, ans, res)
			ans = ans[:len(ans)-1]
			used[i] = false
		}
	}
}

func permute2(nums []int) [][]int {
	result := make([][]int, 0)
	backtrack(nums, len(nums), 0, &result)
	return result
}
func backtrack(nums []int, length, first int, res *[][]int) {
	// 通过两两交换位置的方式来改变
	if first == length {
		*res = append(*res, append([]int{}, nums...))
	}
	for i := first; i < length; i++ {
		nums[i], nums[first] = nums[first], nums[i]
		backtrack(nums, length, first+1, res)
		nums[i], nums[first] = nums[first], nums[i]
	}

}
