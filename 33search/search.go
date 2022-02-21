package main

func main() {

}

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k],
nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如，
[0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
*/

func search2(nums []int, target int) int {
	for i, b := range nums {
		if b == target {
			return i
		}
	}
	return -1
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 左边升序，那就在左边二分,注意，这里left和mid可能是一个值
		if nums[left] <= nums[mid] {
			if nums[left] <= target && nums[mid] >= target {
				right = mid - 1
			} else { // 说明在mid右边部分
				left = mid + 1
			}
		} else { // 右边升序
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
