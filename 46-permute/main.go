package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 4, 6, 2}
	//nums[0], nums[1] = nums[1], nums[0]
	//fmt.Println(nums)

	res := permute2(nums)
	fmt.Println(res)
}
