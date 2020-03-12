package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
	//nums := []int{6, 0, 8, 2, 1, 5}
	res := maxWidthRamp3(nums)
	fmt.Println("res is ", res)
}

/*
给定一个整数数组 A，坡是元组 (i, j)，其中  i < j 且 A[i] <= A[j]。这样的坡的宽度为 j - i。
找出 A 中的坡的最大宽度，如果不存在，返回 0 。
示例 1：
输入：[6,0,8,2,1,5]
输出：4
解释：
最大宽度的坡为 (i, j) = (1, 5): A[1] = 0 且 A[5] = 5.
示例 2：
输入：[9,8,1,0,1,9,4,0,4,1]
输出：7
解释：
最大宽度的坡为 (i, j) = (2, 9): A[2] = 1 且 A[9] = 1.
提示：
    2 <= A.length <= 50000
    0 <= A[i] <= 50000
*/

// 这种方法有解，但是会超出时间限制
func maxWidthRamp(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}

	dp := make(map[int]int)
	// dp[i]的定义是指，在A数据组中，选到下标为i是的最大长度（i可能没有选中）
	max := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if A[j] <= A[i] && i-j > max {
				max = i - j
			}
		}
		dp[i] = max
	}
	return max
}

func maxWidthRamp2(A []int) int {
	if len(A) <= 1 {
		return len(A)
	}

	dp := make(map[int]int)
	// dp[i]的定义是指，在A数据组中，选到下标为i是的最大长度（i可能没有选中）
	max := 0
	for i := 1; i < len(A); i++ {
		if A[i] >= A[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}
	return max
}

// 第二种方法是写单调栈

// 第三种方法，使用排序，并用map记录

func maxWidthRamp3(A []int) int {
	B := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		B[i] = i
	}

	sor := S{
		A: A,
		B: B,
	}
	// 不知道什么原因。这里的排序没有生效
	sort.Sort(sor)
	fmt.Println("sor is ", sor)
	ans := 0
	m := len(A)
	for i := 0; i < len(sor.B); i++ {
		ans = int(math.Max(float64(i-m), float64(m)))
		m = int(math.Min(float64(m), float64(i)))
	}
	return ans
}

type S struct {
	A []int
	B []int
}

func (s S) Len() int {
	return len(s.A)
}

func (s S) Less(i, j int) bool {
	return s.A[i] < s.A[j]
}

func (s S) Swap(i, j int) {
	s.B[i], s.B[j] = s.B[j], s.B[i]
}
