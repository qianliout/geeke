package main

import (
	"fmt"
	"math"
)

func main() {
	res := findKthNumber(13, 2)
	fmt.Println("res is ", res)
}

/*
给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
注意：1 ≤ k ≤ n ≤ 109。
示例 :
输入:
n: 13   k: 2
输出:
10
解释:
字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
*/
func findKthNumber(n int, k int) int {
	prefix, p := 1, 1
	for p < k {
		count := getCount(prefix, n)

		if p+count > k {
			prefix *= 10
			p++
		} else {
			prefix++
			p += count
		}
	}
	return prefix
}

// prefix是前缀，n是上界
func getCount(prefix, n int) int {
	count, cur, nexPrefix := 0, prefix, prefix+1
	for cur <= n {
		count += Min(nexPrefix, n+1) - cur
		cur *= 10
		nexPrefix *= 10
	}
	return count
}

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
