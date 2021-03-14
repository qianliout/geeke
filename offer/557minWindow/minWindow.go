package main

func main() {

}

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：
输入：s = "a", t = "a"
输出："a"
*/
func minWindow(s string, t string) string {
	tt := make(map[byte]int)
	for _, v := range []byte(t) {
		tt[v]++
	}
	exit := make(map[byte]int)
	ss := []byte(s)
	left, right, ans, hasAns := 0, 0, "", false

	for right < len(s) {
		chr := ss[right]
		exit[chr]++
		right++
		for hasAll(exit, tt) {
			hasAns = true
			exit[ss[left]]--
			left++
		}
		if hasAns {
			thisAns := string(ss[left-1 : right])
			if ans == "" || (ans != "" && len(thisAns) < len(ans)) {
				ans = thisAns
			}
		}
	}
	return ans
}

func hasAll(exit, tt map[byte]int) bool {
	for k, v := range tt {
		if exit[k] < v {
			return false
		}
	}
	return true
}
