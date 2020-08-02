package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate3(nums, 3)
	fmt.Println(nums)

}

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例 1:
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:
    尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
    要求使用空间复杂度为 O(1) 的 原地 算法。
*/

func rotate(nums []int, k int) {

}

// 暴力法,反转k次
func Rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
	}
}

// 解法二,把后面k个数存起来,然后把nums移动之后放在前面,但是空间复杂度不对
func Rotate2(nums []int, k int) {
	length := len(nums)
	if k == length {
		return
	}
	if k > length {
		k = k % length
	}

	tem := make([]int, k)

	for i := 0; i < k; i++ {
		tem[i] = nums[length-1-i]
	}
	// 移动Nums

	for i := length - 1; i-k >= 0; i-- {
		nums[i] = nums[i-k]
	}
	// 把临时数组放到前面去
	for i, i2 := range tem {
		nums[k-i-1] = i2
	}
}

// 反转数据
func Rotate3(nums []int, k int) {
	length := len(nums)
	if k%length == 0 {
		return
	}
	k = k % length
	rotateHelper(nums, 0, len(nums)-1)
	rotateHelper(nums, k, length-1)
	rotateHelper(nums, 0, k-1)
}

func rotateHelper(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
