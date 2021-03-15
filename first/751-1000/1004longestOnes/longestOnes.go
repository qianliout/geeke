package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	ans := longestOnes(nums, 3)
	fmt.Println("ans is ", ans)
}

/*
给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
返回仅包含 1 的最长（连续）子数组的长度。
示例 1：
输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：
[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：
输入：A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
提示：
1 <= A.length <= 20000
0 <= K <= A.length
A[i] 为 0 或 1
*/

func longestOnes(A []int, K int) int {
	if K >= len(A) {
		return len(A)
	}
	ans := 0
	left, right := 0, 0
	for right < len(A) {
		if A[right] == 1 {
			right++
		} else {
			if K > 0 {
				right++
				K--
			} else {
				if right-left > ans {
					ans = right - left
				}

				if A[left] == 0 {
					K++
				}
				left++
			}
		}
	}
	if right-left > ans {
		ans = right - left
	}
	return ans
}

func longestOnes2(A []int, K int) int {
	ans, left, right := 0, 0, 0
	for right < len(A) {
		num := A[right]
		right++
		if num != 1 {
			K--
		}
		// 缩小窗口
		for K < 0 {
			if A[left] != 1 {
				K++
			}
			left++
		}
		// 结算结果了
		if right-left > ans {
			ans = right - left
		}
	}

	return ans
}
