package main

import (
	"fmt"
	"math"
)

func main() {
	x := -2312
	res := reverse(x)
	fmt.Println("res is ", res)
}

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
示例 1:
输入: 123
输出: 321
 示例 2:
输入: -123
输出: -321
示例 3:
输入: 120
输出: 21
注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
// 这道题主要是解决溢出的问题
func reverse(x int) int {
	r := x
	num := 0
	for r != 0 {
		i := r % 10
		r = r / 10
		// 检查溢出
		if (num > math.MaxInt32/10) || (num == math.MaxInt32 && i > 7) {
			return 0
		}
		if (num < math.MinInt32/10) || (num == math.MinInt32 && i < -8) {
			return 0
		}
		num = num*10 + i
	}
	return num
}
