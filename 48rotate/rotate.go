package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(input)
	fmt.Println(input)
}
func rotate(maxtrix [][]int) {
	mathRotate(&maxtrix)
}

// 纯数据旋转大法
func mathRotate(matix *[][]int) {
	length := len(*matix)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			(*matix)[i][j], (*matix)[j][i] = (*matix)[j][i], (*matix)[i][j]
		}
	}
	// 上面这个xunh得到如下结构
	//[[1，4，6],[2,5,8],[3,6,9]]
	// 下面这个循环才得到正确的解
	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			(*matix)[i][j], (*matix)[i][length-j-1] = (*matix)[i][length-j-1], (*matix)[i][j]
		}
	}
}
