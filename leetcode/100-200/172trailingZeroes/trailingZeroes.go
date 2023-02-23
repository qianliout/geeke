package main

func main() {

}

/*
给定一个整数 n ，返回 n! 结果中尾随零的数量。
提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
*/

func trailingZeroes(n int) int {
	count := 0

	for n > 0 {
		count += n / 5
		n = n / 5
	}
	return count
}
