package main

func main() {

}

/*
154. 寻找旋转排序数组中的最小值 II,有重重值
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right = right - 1
		}
	}
	return nums[left]
}
