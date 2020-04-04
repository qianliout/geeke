package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
     "123"
     "132"
     "213"
     "231"
     "312"
     "321"
 给定 n 和 k，返回第 k 个排列。
 说明：
     给定 n 的范围是 [1, 9]。
     给定 k 的范围是[1,  n!]。
 示例 1:
 输入: n = 3, k = 3
 输出: "213"
 示例 2:
 输入: n = 4, k = 9
 输出: "2314"
*/

//定义一个全局变量，各个递归中可以更改
var remind int

func main() {
	n := 3
	k := 2
	str := getPermutation2(n, k)
	fmt.Println(str)
}

func getPermutation(n int, k int) string {
	if n == 0 || k == 0 {
		return ""
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	dfs(n, n, k, 0, path, used, &res)
	//fmt.Println("res is ", res)
	result := res[0]
	return strings.Replace(fmt.Sprint(result)[1:len(fmt.Sprint(result))-1], " ", "", -1)
}

// 这种写法是对的，但是会超出时间限制，
func dfs(nums, n, k, left int, path []int, used map[int]bool, res *[][]int) {
	if len(path) == n {
		//fmt.Println("remind1 is", remind, "path is ", path)
		if remind == k-1 {
			*res = append(*res, append([]int{}, path...))
			return
		} else {
			remind += 1
			//fmt.Println("remind2 is", remind)
			return
		}
	}

	for i := 1; i < nums+1; i++ {
		if len(*res) >= 1 {
			break
		}
		u, exit := used[i]
		if exit && u {
			continue
		}

		used[i] = true

		path = append(path, i)
		dfs(nums, n, k, left+1, path, used, res)
		used[i] = false
		path = path[:len(path)-1]
	}
}

var (
	m map[int]int
)

func getPermutation2(n int, k int) string {
	if k == 0 || n == 0 {
		return ""
	}
	nums := make([]int, 0)
	res := make([]int, 0)
	for i := n; i > 0; i-- {
		nums = append(nums, i)
	}
	helper(nums, k, &res)

	ans := ""
	for _, num := range res {
		ans = ans + strconv.Itoa(num)
	}
	return ans
}

func helper(nums []int, k int, res *[]int) {
	if m == nil {
		m = make(map[int]int)
	}
	start := len(nums) - 1

	for start >= 0 {
		j := start

		pre := Fib(len(nums)-1, &m)
		if len(nums) == 1 {
			*res = append(*res, nums[0])
			break
		}
		for k > pre {
			j--
			k -= pre
		}
		*res = append(*res, nums[j])
		nums = append(nums[:j], nums[j+1:]...)
		start = len(nums) - 1
	}
}

func Fib(n int, m *map[int]int) int {
	if n <= 1 {
		return n
	}
	if v, ok := (*m)[n]; ok {
		return v
	}
	r := Fib(n-1, m) * n
	(*m)[n] = r
	return r
}
