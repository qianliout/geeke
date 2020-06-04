package main

func main() {

}

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
示例:
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:
    必须在原数组上操作，不能拷贝额外的数组。
    尽量减少操作次数。
*/

func moveZeroes(nums []int) {
	for len(nums) == 0 {
		return
	}
	i, length := 0, len(nums)-1
	for i <= length {
		if nums[i] == 0 {
			j := i
			for j <= length {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
				j++
			}
		}
		i++
	}
}
