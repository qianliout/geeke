package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := decodeString2("3[a2[c]]")
	fmt.Println("res is ", res)
}

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

type item struct {
	repeat int
	prefix string
}

// 使用栈的思想
func decodeString(s string) string {
	stark := make([]item, 0)
	prefix, repeat := "", 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			stark = append(stark, item{repeat: repeat, prefix: prefix})
			prefix, repeat = "", 0
		} else if s[i] == ']' {
			if len(stark) > 0 {
				last := stark[len(stark)-1]
				stark = stark[:len(stark)-1]
				prefix = last.prefix + strings.Repeat(prefix, last.repeat)
			}
		} else if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			repeat = repeat*10 + num
		} else {
			prefix = prefix + string(s[i])
		}
	}
	return prefix
}

func decodeString2(s string) string {
	res, _ := dfs([]byte(s), 0)
	return res
}

// 使用递归的思想
func dfs(ss []byte, index int) (string, int) {
	prefix, repeat := "", 0

	for index < len(ss) {
		c := ss[index]
		if '0' <= c && '9' >= c {
			n, _ := strconv.Atoi(string(c))
			repeat = repeat*10 + n
		} else if c == '[' {
			tem, i := dfs(ss, index+1)
			prefix += strings.Repeat(tem, repeat)
			index, repeat = i, 0
		} else if c == ']' {
			return prefix, index
		} else {
			prefix = prefix + string(c)
		}
		index++
	}
	return prefix, index
}
