package main

import (
	"strings"
)

func main() {

}

/*
给定两个字符串 a 和 b，寻找重复叠加字符串 a 的最小次数，使得字符串 b 成为叠加后的字符串 a 的子串，如果不存在则返回 -1。
注意：字符串 "abc" 重复叠加 0 次是 ""，重复叠加 1 次是 "abc"，重复叠加 2 次是 "abcabc"。
示例 1：
输入：a = "abcd", b = "cdabcdab"
输出：3
解释：a 重复叠加三遍后为 "abcdabcdabcd", 此时 b 是其子串。
示例 2：
输入：a = "a", b = "aa"
输出：2
示例 3：
输入：a = "a", b = "a"
输出：1
示例 4：
输入：a = "abc", b = "wxyz"
输出：-1
*/

/*
这题最关键的是找到循环退出条件，这是本题的难点。
如果n个A字符串能包含B字符串，可能有几种情况：
1、A = "ab", B = "abab",循环n个A，刚好包含B；
2、A = "ab", B = "ababa",那么需要循环n + 1次；
3、A = "ab", B = "babab"，那么需要循环n + 1次；
4、A = "ab", B = "bababa"，那么需要循环n + 2次；
如果B不满足以上情况，A再怎么循环也是白搭，比如，A = "ab", B = "bababb"，循环多少次都只是浪费。
因此，如果B能够在A的N次循环中被找到，最多只需要循环n+2次。循环终止条件就是累积的字符串长度 > (n + 2) * sizeA = (sizeB/sizeA + 2) * sizeA= sizeB + 2*sizeA.
另外，这里求最小次数的含义体现在，对于满足条件是A和B，有可能循环n次，n + 1次，或n + 2次，就包含B，再循环更多次数当然也能找到，当包含后求得的次数就是最小次数。
而对于无法通过循环包含B的A，我们要找到一个最多遍历的次数，恰好就是上面说的最小次数，因为对于这种情况，循环再多也找不到。
*/

// https://leetcode-cn.com/problems/repeated-string-match/solution/golang-zui-da-chang-du-2ab-by-bloodborne/
func repeatedStringMatch(a string, b string) int {
	la, lb := len(a), len(b)
	maxLen := la*2 + lb
	t, count := "", 0
	for len(t) < maxLen {
		if strings.Contains(t, b) {
			return count
		}
		t += a
		count++
	}
	return -1
}
