package main

import (
	"fmt"
)

func main() {
	ans := findDuplicate([]int{1, 2, 2, 2, 2})
	fmt.Println(ans)
}

func findDuplicate(nums []int) int {
	ans := 0
	for i, v := range nums {
		ans = ans ^ i ^ v
	}
	return ans
}
