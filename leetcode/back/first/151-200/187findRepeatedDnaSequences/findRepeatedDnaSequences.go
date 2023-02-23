package main

import (
	"fmt"
)

func main() {
	s := "AAAAAAAAAAA"
	res := findRepeatedDnaSequences(s)
	fmt.Println("res is ", res)
}

/*
所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA
中的重复序列有时会对研究非常有帮助。
编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。
示例：
输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]
*/
func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)
	if len(s) < 10 {
		return res
	}
	exit := make(map[string]int)
	ss := []byte(s)
	for i := 0; i+10 <= len(ss); i++ {
		exit[string(ss[i:i+10])] += 1
	}
	for word, i := range exit {
		if i > 1 {
			res = append(res, word)
		}
	}
	fmt.Println(exit)
	return res
}
