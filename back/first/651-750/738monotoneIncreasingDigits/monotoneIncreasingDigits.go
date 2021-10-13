package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(monotoneIncreasingDigits(123))
}

/*
给定一个非负整数 N，找出小于或等于 N 的最大的整数，同时这个整数需要满足其各个位数上的数字是单调递增。
（当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。）
示例 1:
输入: N = 10
输出: 9
示例 2:
输入: N = 1234
输出: 1234
示例 3:
输入: N = 332
输出: 299
*/

// https://leetcode-cn.com/problems/monotone-increasing-digits/solution/jian-dan-tan-xin-shou-ba-shou-jiao-xue-k-a0mp/
func monotoneIncreasingDigits(N int) int {
	ss := []byte(strconv.Itoa(N))
	start, max := -1, -1
	for i := 0; i < len(ss)-1; i++ {
		if int(ss[i]) > max {
			max = int(ss[i])
			start = i
		}
		if ss[i] > ss[i+1] {
			ss[start]--
			for j := start + 1; j < len(ss); j++ {
				ss[j] = '9'
			}
			break
		}
	}
	a, _ := strconv.Atoi(string(ss))
	return a
}
