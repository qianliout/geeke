package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含非负整数的数组，你的任务是统计其中可以组成三角形三条边的三元组个数。
示例 1:
输入: [2,2,3,4]
输出: 3
解释:
有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3
注意:
数组长度不超过1000。
数组里整数的范围为 [0, 1000]。
*/
// [0,1,1,1] 得到结果是1，这里就不应该,结果 应该是0
func main() {
	nums := []int{0, 1, 1, 1}
	res := triangleNumber(nums)
	fmt.Println("res is ", res)
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	s := make([]int, max+1)
	for _, n := range nums {
		s[n]++
	}
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := nums[j] + 1; k < nums[j]+nums[i]; k++ {
				if k > 0 && k <= max {
					ans += s[k]
				}
			}

		}
	}
	return ans
}
