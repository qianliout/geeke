package main

import (
	"fmt"
	"math"
)

func main() {
	ans := longestCommonPrefix([]string{"dog", "racecar", "car"})
	fmt.Println("ans is ", ans)
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"
*/

func longestCommonPrefix(strs []string) string {
	min, minLen := "", math.MaxInt64
	for i := range strs {
		if len(strs[i]) < minLen {
			min, minLen = strs[i], len(strs[i])
		}
	}
	length, finish := 0, true
	for length < minLen {
		for j := range strs {
			if strs[j][length] != min[length] {
				finish = false
				break
			}
		}
		if !finish {
			break
		}
		length++
	}
	return string([]byte(min)[:length])
}
