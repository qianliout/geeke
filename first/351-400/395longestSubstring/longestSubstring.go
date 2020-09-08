package main

import (
	"fmt"
	"math"
)

func main() {
	res := longestSubstring("ababbc", 2)
	fmt.Println("res is ", res)
}

/*
找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。
示例 1:
输入:
s = "aaabb", k = 3
输出:
3
最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2:
输入:
s = "ababbc", k = 2
输出:
5   // 怎么感觉这里有问题呢，应该是4，b重复两次就行了啊
最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。
*/

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	ss := []byte(s)
	left, cur, min := 0, 0, math.MaxInt64
	used := make(map[byte]int)
	for cur < len(s) {
		c := ss[cur]
		used[c]++
		cur++
		for isTrue(used, k) {
			if cur-left < min {
				min = cur - left
			}

			lc := ss[left]
			used[lc]--
			left++
		}

	}
	return min
}

func isTrue(m map[byte]int, k int) bool {
	for _, l := range m {
		if l > 0 && l < k {
			return false
		}
	}
	return true
}
