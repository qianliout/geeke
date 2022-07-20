package main

import (
	"fmt"
)

func main() {
	sum3 := combinationSum3(3, 9)
	fmt.Println("sum3  is ", sum3)
}

/*
找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
只使用数字1到9
每个数字 最多使用一次 
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
*/

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	if k > n || k*9 < n {
		return ans
	}
	dsf(k, n, 1, []int{}, &ans)
	return ans
}

func dsf(k, n, start int, path []int, ans *[][]int) {
	if n == 0 && len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if n < 0 || len(path) > k {
		return
	}
	for i := start; i <= 9; i++ {
		// if len(path) > 0 && i <= path[len(path)-1] {
		// 	continue
		// }
		// path = append(path, i)
		// n = n - i
		dsf(k, n-i, i+1, append(path, i), ans)
		// n = n + i
		// path = path[:len(path)-1]
	}
}
