package main

func main() {

}

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。
*/
func generate(numRows int) [][]int {
	res := make([][]int, 0)
	if numRows == 0 {
		return res
	}
	first := []int{1}
	res = append(res, first)

	for i := 2; i <= numRows; i++ {
		second := helper(first)
		res = append(res, second)
	}
	return res
}

func helper(src []int) []int {
	res := make([]int, 0)
	if len(src) == 0 {
		return res
	}
	n := len(src)
	for i := 0; i <= n; i++ {
		if i-1 < 0 {
			res = append(res, src[i])
		}
		if i == n {
			res = append(res, src[i-1])
		}
		if i-1 >= 0 && i != n {
			res = append(res, src[i-1]+src[i])
		}
	}
	return res
}
