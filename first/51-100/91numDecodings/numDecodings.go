package main

import (
	"fmt"
)

func main() {
	res := numDecodings("226")
	fmt.Println("res is ", res)
}

/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：
'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。
示例 1:
输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:
输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
*/
// 这道题当然可以不使用[]int的方式，但是这里我不想再处理了
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	bs := []byte(s)

	nums := make([]int, 0)
	for _, i2 := range bs {
		nums = append(nums, int(i2)-48)
	}
	return numdecode(nums)
}

func numdecode(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make(map[int]int)

	if nums[0] == 0 {
		return 0
	}

	dp[0], dp[-1] = 1, 1 // dp[-1]=1是这道题的关键

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[i-1] == 1 || nums[i-1] == 2 {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else if nums[i-1] == 1 || (nums[i-1] == 2 && nums[i] >= 1 && nums[i] <= 6) {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	//fmt.Println(dp)
	return dp[len(nums)-1]
}
