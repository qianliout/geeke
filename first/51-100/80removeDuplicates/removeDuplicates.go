package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//nums := []int{1, 1, 1, 2, 2, 3}
	res := removeDuplicates(nums)
	fmt.Println("res is ", res)
}

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
示例 1:
给定 nums = [1,1,1,2,2,3],
函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
你不需要考虑数组中超出新长度后面的元素。
示例 2:
给定 nums = [0,0,1,1,1,1,2,3,3],
函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。
你不需要考虑数组中超出新长度后面的元素。
*/
func removeDuplicates(nums []int) int {
	res := removeDuplicatesForK(nums, 2)
	return res
}

func removeDuplicatesForK(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 {
		return len(nums)
	}
	index, count := 0, 1

	for left := 1; left < len(nums); left++ {
		if nums[left] == nums[index] {
			if count < k {
				index++
				count++
				nums[index] = nums[left]
			}
		} else {
			index++
			count = 1
			nums[index] = nums[left]
		}

	}
	//fmt.Println(nums[:index+1])
	return index + 1
}
