package main

import (
	"fmt"
)

func main() {
	nums := []int{0}
	res := sortColors2(nums)
	fmt.Println(res)
}

/*
75. 颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
注意:
不能使用代码库中的排序函数来解决这道题。
示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
进阶：
    一个直观的解决方案是使用计数排序的两趟扫描算法。
    首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
    你能想出一个仅使用常数空间的一趟扫描算法吗？
*/
// 这种方法很通用，可以使用多数的排序，但是题目中说了只有三种数据，所以可以有更好的方法
func sortColors(nums []int) []int {
	if len(nums) == 1 || len(nums) == 0 {
		for i := 1; i < len(nums); i++ {
			for k := i; k > 0; k-- {
				if nums[k] < nums[k-1] {
					nums[k], nums[k-1] = nums[k-1], nums[k]
					if nums[k] == nums[k-1] {
						break
					}
				}
			}
		}
	}
	return nums
}

// 因为只有三个数字，所以我们知道，0一定是在最前面，2，一定是在最后边，所以可以使用双指针的方式
func sortColors2(nums []int) []int {
	if len(nums) > 1 {
		p1, curr := 0, 0
		p2 := len(nums) - 1
		for curr <= p2 {
			currentNum := nums[curr]
			if currentNum == 2 {
				nums[curr], nums[p2] = nums[p2], nums[curr]
				p2--
			} else if currentNum == 0 {
				nums[curr], nums[p1] = nums[p1], nums[curr]

				p1++
				curr++
			} else {
				curr++
			}
		}
	}
	return nums
}
