package main

func main() {
	res := []int{2, 2, 0, 2, 1, 1, 0}
	sortColors(res)
}

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
*/
func sortColors(nums []int) {
	red, white, cur := 0, len(nums)-1, 0
	for cur < len(nums) {
		if nums[cur] == 2 {
			nums[cur], nums[white] = nums[white], nums[cur]
			white--
		} else if nums[cur] == 0 {
			nums[cur], nums[red] = nums[red], nums[cur]
			red++
			cur++
		} else {
			cur++
		}
	}
}

func sortColors2(nums []int) {
}
