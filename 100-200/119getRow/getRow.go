package main

import (
	"fmt"
)

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(3))
}

/*
给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/
func getRow(rowIndex int) []int {
	first := []int{1}
	for i := 0; i < rowIndex; i++ {
		first = helper(first)
	}
	return first
}

func helper(data []int) []int {
	res := make([]int, 0)
	if len(data) == 0 {
		return res
	}
	for i := 0; i <= len(data); i++ {
		if i-1 < 0 {
			res = append(res, data[i])
		} else if i == len(data) {
			res = append(res, data[i-1])
		} else {
			res = append(res, data[i-1]+data[i])
		}
	}
	return res
}
