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

func getRow(rowIndex int) []int {
	kRows := make([]int, rowIndex+1)
	for k := 0; k <= rowIndex; k++ {
		kRows[k] = 1
		for j := k; j > 1; j-- {
			kRows[j-1] = kRows[j-2] + kRows[j-1]
		}
	}
	return kRows

}
