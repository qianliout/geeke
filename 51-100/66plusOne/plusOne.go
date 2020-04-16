package main

import (
	"fmt"
)

func main() {
	res := plusOne([]int{1})
	fmt.Println("res is ", res)

}

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
示例 1:
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {
	n := len(digits)
	res := make([]int, n+1)
	if n == 0 {
		return res
	}
	for i := n - 1; i >= 0; i-- {
		num := digits[i]
		if i == n-1 {
			num += 1
		}
		num = num + res[i+1]
		if num >= 10 {
			res[i+1] = num - 10
			res[i] += 1
		} else {
			res[i+1] = num
		}
	}
	if res[0] == 0 {
		return res[1:]
	}
	return res
}
