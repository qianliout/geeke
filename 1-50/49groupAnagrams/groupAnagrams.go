package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams3(input)
	fmt.Println(res)
}

//字母异位词分组  排序方法
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}
	var helpMap = make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		var curStrArr = strings.Split(strs[i], "")
		sort.Strings(curStrArr)
		var curStrIndex = strings.Join(curStrArr, "")
		helpMap[curStrIndex] = append(helpMap[curStrIndex], strs[i])
	}
	for _, value := range helpMap {
		res = append(res, value)
	}
	return res
}

//在go语言中数组（不是slice）是可以hash的，也就是说，是可以直接作为map的key使用

func groupAnagrams2(strs []string) [][]string {
	var res [][]string
	if len(strs) == 0 {
		return res
	}
	var helpMap = make(map[[26]int][]string)
	for i, str := range strs {
		var strArray [26]int
		for _, b := range []byte(str) {
			strArray['z'-b] += 1
		}
		helpMap[strArray] = append(helpMap[strArray], strs[i])
	}
	for _, value := range helpMap {
		res = append(res, value)
	}
	return res
}

func groupAnagrams3(strs []string) [][]string {
	m := map[string][]string{}
	r := [][]string{}

	for i := 0; i < len(strs); i++ {
		h := map[byte]int{}

		for j := 0; j < len(strs[i]); j++ {
			if v, ok := h[strs[i][j]]; ok {
				h[strs[i][j]] = v + 1
			} else {
				h[strs[i][j]] = 1
			}
		}
		k := fmt.Sprint(h) // 这个函数就是把其他结构表示成字符串形式，类型python的str()方法
		if v, ok := m[k]; ok {
			m[k] = append(v, strs[i])
		} else {
			m[k] = []string{strs[i]}
		}
	}

	for _, v := range m {
		r = append(r, v)
	}

	return r
}
