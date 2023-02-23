package main

func main() {

}

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/
func longestConsecutive2(nums []int) int {
	numMap := make(map[int]int, len(nums))
	for i := range nums {
		numMap[nums[i]]++
	}
	ans := 0

	for _, num := range nums {
		if numMap[num-1] <= 0 {
			cru, this := num, 0
			for numMap[cru] > 0 {
				cru++
				this++
			}
			if this > ans {
				ans = this
			}
		}
	}
	return ans
}
