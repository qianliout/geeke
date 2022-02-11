package main

import (
	"fmt"
	"math"
)

func main() {
	x := reverse3(-123)
	fmt.Println(x)
}

/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
示例 1：
输入：x = 123
输出：321
示例 2：
输入：x = -123
输出：-321
*/

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	positive := true
	if x < 0 {
		positive = false
		x = -x
	}
	res := 0
	for x > 0 {
		res *= 10
		res = res + (x % 10)
		x /= 10
	}
	if !positive {
		res = -res
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}

func reverse3(x int) int {
	if x == 0 {
		return 0
	}
	res := 0
	for x != 0 {
		res = res*10 + (x % 10)
		x /= 10
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}

func reverse2(x int) int {
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
