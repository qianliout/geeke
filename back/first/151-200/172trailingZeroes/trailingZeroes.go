package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(25))
	res := trailingZeroes(30)
	fmt.Println("res is ", res)
}

/*
给定一个整数 n，返回 n! 结果尾数中零的数量。
示例 1:
输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:
输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
*/

// 2*5后面有个0,只要有5那么一定就会有一个2,所以只有求出有多少个5就行
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n = n / 5 // 解决25的这情况，因为25有两个5的因子
	}
	return count
}
