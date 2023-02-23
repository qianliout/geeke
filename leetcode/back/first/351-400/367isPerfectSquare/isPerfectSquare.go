package main

import (
	"fmt"
)

func main() {
	res := isPerfectSquare2(16)
	fmt.Println("res is ", res)
}

/*
给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
说明：不要使用任何内置的库函数，如  sqrt。
示例 1：
输入：16
输出：True
示例 2：
输入：14
输出：False
*/
// 这道题最好的方式是打表法,因为就那么几个数
func isPerfectSquare(num int) bool {
	exit := make(map[int]int)
	three := 0
	for i := 1; three <= num; i++ {
		three = perfect(i)
		exit[three] = i
	}

	fmt.Println(exit)
	_, ok := exit[num]
	if ok {
		return true
	}
	return false
}

func perfect(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * n
	}
	return res
}

func isPerfectSquare2(num int) bool {
	if num <= 0 {
		return false
	}
	subtrahend := 1
	for num > 0 {
		num -= subtrahend
		subtrahend += 2
	}
	if num == 0 {
		return true
	}
	return false
}
