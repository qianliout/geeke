package main

import (
	"strings"
)

func main() {

}

/*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
文本的最后一行应为左对齐，且单词之间不插入额外的空格。
注意:
单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。
*/
func fullJustify(words []string, maxWidth int) (ans []string) {
	i, n := 0, len(words)
	for {
		start := i   // 这一行文本在 words 的起始位置
		wordLen := 0 // 统计这一行单词字符的个数
		for ; i < n && wordLen+len(words[i])+i-start <= maxWidth; i++ {
			wordLen += len(words[i])
		}
		// 循环结束后，i 为这一行文本在 words 的结束位置 +1
		if i == n { // 最后一行
			s := strings.Join(words[start:], " ")
			ans = append(ans, s+strings.Repeat(" ", maxWidth-len(s)))
			return
		}
		space := maxWidth - wordLen // 需要插入至单词间的空格个数
		if i-start == 1 {           // 只有一个单词的情况，在末尾加上空格即可
			ans = append(ans, words[start]+strings.Repeat(" ", space))
		} else { // 计算两单词间至少要加多少个空格，以及有多少个间隔要额外加一个空格
			avg, extra := space/(i-start-1), space%(i-start-1)
			avgSpace := strings.Repeat(" ", avg)
			s1 := strings.Join(words[start:start+extra+1], avgSpace+" ")
			s2 := strings.Join(words[start+extra+1:i], avgSpace)
			ans = append(ans, s1+avgSpace+s2)
		}
	}
}
