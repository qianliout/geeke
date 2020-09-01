package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := isAdditiveNumber("112358")
	fmt.Println("res is ", res)

}

/*
累加数是一个字符串，组成它的数字可以形成累加序列。
一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。
给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。
说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。
示例 1:
输入: "112358"
输出: true
解释: 累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
示例 2:
输入: "199100199"
输出: true
解释: 累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
进阶:
你如何处理一个溢出的过大的整数输入?
*/
func isAdditiveNumber(num string) bool {
	var res bool
	backTrack([]byte(num), 0, 0, 0, &res)
	return res
}

func backTrack(nums []byte, usedCount int, first, second int, res *bool) {
	if usedCount <= 2 && len(nums) == 0 {
		*res = false
		return
	}
	if len(nums) == 0 {
		*res = true
		return
	}

	for i := 0; i < len(nums); i++ {
		// 剪枝--单个的0可以作为一个数，而01，02之类的不能，直接返回
		n, _ := strconv.Atoi(string(nums[:i+1]))
		if len(strconv.Itoa(n)) != len(string(nums[:i+1])) {
			return
		}
		if usedCount >= 2 {
			// 当前i和前两个数符合
			if n == first+second {
				backTrack(nums[i+1:], usedCount+1, second, n, res)
			}
		} else {
			// 前面还没有两个数时，就直接回漱验证
			backTrack(nums[i+1:], usedCount+1, second, n, res)
		}

		if *res {
			return
		}
	}
}
