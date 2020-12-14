package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeSquareSum2(0))
}

/*
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
示例 1：
输入：c = 5
输出：true
解释：1 * 1 + 2 * 2 = 5
*/
// 方法正确，但是会超时,没有记忆化
func judgeSquareSum(c int) bool {
	// 0的平方是0，所以这里只能是小于等于
	for i := 0; i <= c; i++ {
		for j := 0; i*i+j*j <= c; j++ {
			if i*i+j*j == c {
				return true
			}
		}
	}
	return false
}

func judgeSquareSum2(c int) bool {
	mem := make(map[int]bool)
	// 第一步，加数
	for i := 0; i*i <= c; i++ {
		mem[i*i] = true
	}
	// 第二步检测
	for i := 0; i*i <= c; i++ {
		if mem[c-i*i] {
			return true
		}
	}
	return false
}
