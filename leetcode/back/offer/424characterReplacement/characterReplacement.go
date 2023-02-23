package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
注意：字符串长度 和 k 不会超过 104。
示例 1：
输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。
示例 2：
输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
*/

func characterReplacement(s string, k int) int {
	ss := []byte(s)
	ans, left, right, max := 0, 0, 0, 0
	widow := make(map[byte]int)
	for right < len(s) {
		chr := ss[right]
		right++
		widow[chr]++
		if widow[chr] > max {
			max = widow[chr]
		}
		for right-left > max+k {
			widow[ss[left]]--
			left++
		}
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

func characterReplacement2(s string, k int) int {
	length := len(s)
	if length <= 0 {
		return 0
	}
	ss := []byte(s)
	frequencyMap := make(map[byte]int)
	left, right, charMax := 0, 0, 0
	ans := 0
	for left < length && right < length {
		cha := ss[right]
		frequencyMap[cha]++
		charMax = Max(charMax, frequencyMap[cha])
		right++
		// 缩小windown
		if right-left > charMax+k {
			frequencyMap[ss[left]]--
			left++
		}
		// 更新结果
		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
