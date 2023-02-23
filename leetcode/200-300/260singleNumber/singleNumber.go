package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	number := singleNumber(nums)
	fmt.Println(number)
}
func singleNumber(nums []int) []int {
	flag := 0
	for i := range nums {
		flag = flag ^ nums[i]
	}
	diff := flag & (-flag) // 最后一个1

	left, right := make([]int, 0), make([]int, 0)
	for i := range nums {
		if nums[i]&diff == 0 {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	ans := make([]int, 0)
	if len(left) > 0 {
		tem := 0
		for i := range left {
			tem = tem ^ left[i]
		}
		ans = append(ans, tem)
	}

	if len(right) > 0 {
		tem := 0
		for i := range right {
			tem = tem ^ right[i]
		}
		ans = append(ans, tem)
	}
	return ans
}
