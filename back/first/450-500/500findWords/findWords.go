package main

import (
	"fmt"
)

func main() {
	fmt.Println('A')
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	fmt.Println("res is ", res)
}

/*
给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
American keyboard
示例：
输入: ["Hello", "Alaska", "Dad", "Peace"]
输出: ["Alaska", "Dad"]
注意：
    你可以重复使用键盘上同一字符。
    你可以假设输入的字符串将只包含字母
*/

func findWords(words []string) []string {
	res := make([]string, 0)
	if len(words) == 0 {
		return res
	}
	hash := make(map[uint8]int)
	for _, ch := range []byte("qwertyuiop") {
		hash[ch] = 1
		hash[ch-32] = 1
	}

	for _, ch := range []byte("asdfghjkl") {
		hash[ch] = 2
		hash[ch-32] = 2
	}

	for _, ch := range []byte("zxcvbnm") {
		hash[ch] = 3
		hash[ch-32] = 3
	}

	fmt.Println(hash)
	for _, w := range words {
		add := true
		// ww := []byte(w)
		for i := 0; i < len(w)-1; i++ {
			if hash[w[i]] != hash[w[i+1]] {
				add = false
				break
			}
		}
		if add {
			res = append(res, w)
		}
	}
	return res
}
