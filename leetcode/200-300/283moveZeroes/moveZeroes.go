package main

func main() {
	nums := []int{0, 0, 1, 0, 3, 1, 2, 0, 0, 12, 0}
	moveZeroes(nums)
}

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
func moveZeroes(nums []int) {
	left := 0
	for left < len(nums) {
		if nums[left] != 0 {
			left++
			continue
		}
		// 找下一个可以替换的非0
		right := left + 1
		for right < len(nums) {
			if nums[right] == 0 {
				right++
			} else {
				break
			}
		}
		// 找到最后都没有找到，说明再没有非0的值了
		if right >= len(nums) {
			return
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
}
