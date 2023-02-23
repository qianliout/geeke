package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println("res is ", res)
}

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

func groupAnagrams1(strs []string) [][]string {
	ss := make([]string, 0)

	for i := range strs {
		s := []byte(strs[i])
		sort.Sort(Bits(s))
		ss = append(ss, string(s))
	}

	resMap := make(map[string][]int)

	for i, v := range ss {
		if resMap[v] == nil {
			resMap[v] = make([]int, 0)
		}
		resMap[v] = append(resMap[v], i)
	}
	res := make([][]string, 0)
	for _, v := range resMap {
		vv := make([]string, 0)
		for _, i := range v {
			vv = append(vv, strs[i])
		}
		res = append(res, vv)
	}
	return res
}

type Bits []byte

func (bs Bits) Len() int {
	return len(bs)
}

func (bs Bits) Less(i, j int) bool {
	return bs[i] < bs[j]
}

func (bs Bits) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	sortMap := make(map[string][]string)

	for _, str := range strs {
		ss := strings.Split(str, "")
		sort.Strings(ss)
		key := strings.Join(ss, "")
		sortMap[key] = append(sortMap[key], str)
	}
	for _, v := range sortMap {
		ans = append(ans, v)
	}
	return ans
}
