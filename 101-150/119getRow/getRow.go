package main

import (
	"fmt"
)

func main() {
	res := getRow(3)
	fmt.Println("res is ", res)
}

/*
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
示例:
输入: 3
输出: [1,3,3,1]
进阶：
你可以优化你的算法到 O(k) 空间复杂度吗？
*/

// 第一种写法和118相同，这里就不再写

var Cnmap map[int]int

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	// 奇数
	if rowIndex&1 == 1 {
		for i := 0; i <= rowIndex/2; i++ {
			res[i] = Cn(i + 1)
			res[rowIndex-i] = Cn(i + 1)
		}
	} else {
		for i := 0; i < rowIndex/2; i++ {
			res[i] = Cn(i + 1)
			res[rowIndex-i] = Cn(i + 1)
		}
		res[rowIndex/2] = Cn(rowIndex + 1)
	}
	return res
}

func Cn(n int) int {
	if Cnmap == nil {
		Cnmap = make(map[int]int)
	}
	if Cnmap[n] > 0 {
		return Cnmap[n]
	}
	if n <= 1 {
		return n
	}
	r := n * Cn(n-1)
	Cnmap[n] = r
	return r
}
