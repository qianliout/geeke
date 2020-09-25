package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	res := findAllConcatenatedWordsInADict(words)
	fmt.Println("res is ", res)
}

/*
给定一个不含重复单词的列表，编写一个程序，返回给定单词列表中所有的连接词。
连接词的定义为：一个字符串完全是由至少两个给定数组中的单词组成的。
示例:
输入: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
输出: ["catsdogcats","dogcatsdog","ratcatdogcat"]
解释: "catsdogcats"由"cats", "dog" 和 "cats"组成;
     "dogcatsdog"由"dog", "cats"和"dog"组成;
     "ratcatdogcat"由"rat", "cat", "dog"和"cat"组成。
说明:
    给定数组的元素总数不超过 10000。
    给定数组中元素的长度总和不超过 600000。
    所有输入字符串只包含小写字母。
    不需要考虑答案输出的顺序。
*/

/*
words按照长度从低向高排序
第一个单词长度最短肯定是不是结果
不是返回结果的单词 加入set中 认为是一个普通的单词
然后从第二个单词开始 每次取单词的一个字母加上单词中之前的字符串 判断set中是否存在
如果存在 并且到达单词末尾 则找到
如果存在 但是未到达末尾 继续寻找 将起始字符串的索引加1
否则失败 将单词加入set中
*/

func findAllConcatenatedWordsInADict(words []string) []string {
	res := make([]string, 0)
	set := make(map[string]bool)

	if len(words) <= 1 {
		return res
	}

	sort.Sort(Item(words))
	fmt.Println("word is ", words)
	set[words[0]] = true

	for i := 1; i < len(words); i++ {
		if getWord([]byte(words[i]), 0, &set) {
			res = append(res, words[i])
		} else {
			set[words[i]] = true
		}
	}
	return res
}

func getWord(word []byte, start int, set *map[string]bool) bool {
	length := len(word)
	str := ""
	for start < length {
		str = str + string([]byte{word[start]})
		if (*set)[str] && (start == length-1 || getWord(word, start+1, set)) {
			return true
		}
		start++
	}
	return false
}

type Item []string

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	return len(it[i]) < len(it[j])
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
