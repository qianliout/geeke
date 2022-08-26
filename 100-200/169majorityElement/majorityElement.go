package main

func main() {

}

/*
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素
*/

// 摩尔投票法
func majorityElement(nums []int) int {
	candidate, cnt := nums[0], 0

	for i := range nums {
		if cnt == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			cnt++
		} else {
			cnt--
		}
	}
	return candidate
}
