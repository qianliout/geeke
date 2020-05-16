package main

import (
	. "outback/leetcode/common"
)

func main() {

}

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换 k 次。
在执行上述操作后，找到包含重复字母的最长子串的长度。
注意:
字符串长度 和 k 不会超过 104。
示例 1:
输入:
s = "ABAB", k = 2
输出:
4
解释:
用两个'A'替换为两个'B',反之亦然。
示例 2:
输入:
s = "AABABBA", k = 1
输出:
4
解释:
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
*/

// 题解:https://leetcode-cn.com/problems/longest-repeating-character-replacement/solution/tong-guo-ci-ti-liao-jie-yi-xia-shi-yao-shi-hua-don/
func characterReplacement(s string, k int) int {
	length := len(s)
	if length <= 0 {
		return 0
	}
	ss := []byte(s)
	frequencyMap := make(map[byte]int)
	left, right, charMax := 0, 0, 0
	for left < length && right < length {
		index := ss[right] - 'A'
		frequencyMap[index]++
		charMax = Max(charMax, frequencyMap[index])
		if right-left+1 > charMax+k {
			frequencyMap[ss[left]-'A']--
			left++
		}
		right++
	}
	return length - left
}
