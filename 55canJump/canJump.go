package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

//用current记录还能往前进多少格，每次进一current则为current-1和当前所处位置的值中的较大值。
//当current为0时，代表无法前进，返回False。
//反之则为True
func canJump(nums []int) bool {
	current := 0
	for i := 0; i < len(nums)-1; i++ {
		current = int(math.Max(float64(current-1), float64(nums[i])))
		if current == 0 {
			return false
		}
	}
	return true
}
