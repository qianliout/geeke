package main

import (
	"fmt"
)

func main() {
	res := countDigitOne(1234) //12
	fmt.Println("res is ", res)
}

/*
给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
示例:
输入: 13
输出: 6
解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。
*/
func countDigitOne(n int) int {
	// 先算个位的
	//xyzdabc
	//依次考虑个位、十位、百位...是 1
	//k = 1000, 对应于上边举的例子
	count := 0
	for i := 1; i <= n; i = i * 10 {
		abc := n % i
		xyzd := n / i
		d := xyzd % 10
		xyz := xyzd / 10

		if d == 0 {
			count = count + xyz*i
		} else if d == 1 {
			count = count + xyz*i + abc + 1
		} else if d > 1 {
			count = count + xyz*i + i
		}

		//	防止溢出
		if xyz == 0 {
			break
		}
	}
	return count
}
