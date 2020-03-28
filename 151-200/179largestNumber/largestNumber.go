package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("ehllo ", "" > "30")

	nums := []int{2, 0}
	//nums := []int{10, 2}
	//nums := []int{128, 12}
	res := largestNumber(nums)
	fmt.Println("res is ", res)
}

/*
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
示例 1:
输入: [10,2]
输出: 210
示例 2:
输入: [3,30,34,5,9]
输出: 9534330
说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
链接：https://leetcode-cn.com/problems/largest-number
*/
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	var s Numlist
	for i := 0; i < len(nums); i++ {
		numstr := strconv.Itoa(nums[i])
		s = append(s, numstr)
	}
	sort.Sort(s)
	//fmt.Println(s)
	if s[0] == "0" {
		return "0"
	}
	var ans string
	for i := 0; i < len(s); i++ {
		ans += s[i]
	}
	return ans
}

type Numlist []string

func (n Numlist) Len() int      { return len(n) }
func (n Numlist) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Numlist) Less(i, j int) bool {
	fmt.Printf("od %T,%+v\n", n[i]+n[j], n[i]+n[j])
	return n[i]+n[j] > n[j]+n[i]
}
