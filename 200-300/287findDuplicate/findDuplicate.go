package main

import (
	"fmt"
)

func main() {
	ans := findDuplicate([]int{1, 2, 2, 2, 2})
	fmt.Println(ans)
}

// 确定有环
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow, fast = nums[slow], nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	fast = 0
	for fast != slow {
		fast, slow = nums[fast], nums[slow]
	}
	return fast
}
