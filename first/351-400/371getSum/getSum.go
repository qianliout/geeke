package main

func main() {

}

/*
不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
示例 1:
输入: a = 1, b = 2
输出: 3
示例 2:
输入: a = -2, b = 3
输出: 1
*/

// 位运算
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	low, high := 0, 0
	for {
		low = a ^ b
		high = a & b
		if high == 0 {
			break
		}
		a = low
		b = high << 1
	}
	return low
}
