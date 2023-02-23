package main

func main() {

}

/*
给定字符串 S 和单词字典 words, 求 words[i] 中是 S 的子序列的单词个数。
示例:
输入:
S = "abcde"
words = ["a", "bb", "acd", "ace"]
输出: 3
解释: 有三个是 S 的子序列的单词: "a", "acd", "ace"。
注意:
所有在words和 S 里的单词都只由小写字母组成。
S 的长度在 [1, 50000]。
words 的长度在 [1, 5000]。
words[i]的长度在[1, 50]
*/
// 子序列可以不连序  "abcde" ["a","bb","acd","ace"] 结果是3

func numMatchingSubseq(S string, words []string) int {
	ans := 0
	used := make(map[string]int)
	// 给定的words可能有重
	for _, w := range words {
		used[w]++
	}
	for k, w := range used {
		if check(S, k) {
			ans += w
		}
	}
	return ans
}

// 这种方法一定要掌握
func check(pre, sub string) bool {
	i, j := 0, 0
	for i < len(pre) && j < len(sub) {
		if pre[i] == sub[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == len(sub)
}
