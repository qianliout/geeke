package main

import (
	"fmt"
)

func main() {
	and := rangeBitwiseAnd(2147483645, 2147483647)
	fmt.Println("and ", and)
}

/*
给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）
*/
func rangeBitwiseAnd(left int, right int) int {
	if left > right {
		left, right = right, left
	}
	move := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		move++
	}
	return left << move
}
