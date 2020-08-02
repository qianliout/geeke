package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := start(3)
	//res := findNthDigit(365)
	fmt.Println("res is ", res)
}

/*
在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。

注意:
n 是正数且在32位整数范围内 ( n < 231)。
示例 1:
输入:
3
输出:
3
示例 2:
输入:
11
输出:
0
说明:
第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。
*/
func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	dight := n
	// 第一步，确定落在那一组
	j := 1
	i := 1

	for dight > 9*j*i {
		dight = dight - (9*j)*i
		j = j * 10
		i++
	}
	// 再确定是那个值
	r := ((n - start(i)) / i) + start(i)
	// 再确定那个值下的index
	idx := (n - start(i)) % i
	if idx == 0 {
		idx = i - 1
	} else {
		idx = idx - 1
	}
	fmt.Println("i is ", i)
	fmt.Println("r is ", r)
	fmt.Println("index is ", i)
	fmt.Println("start i ", start(i))

	res := string([]byte(strconv.Itoa(r))[idx])
	atoi, _ := strconv.Atoi(res)
	return atoi
}

func start(n int) int {
	if n < 2 {
		return 0
	}
	res := 9
	for i := 0; i < n-1; i++ {
		res += res * 10 * i
	}
	return res
}
