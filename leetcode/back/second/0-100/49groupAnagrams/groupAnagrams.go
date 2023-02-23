package main

import (
	"fmt"
)

func main() {
	strs := []string{"rod", "dye"}
	res := groupAnagrams(strs)
	fmt.Println("res is ", res)
}

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
*/

// 取质数，用hash的思想,这里只能用乘法，不能用加法
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[int][]string)
	res := make([][]string, 0)
	primeNumber := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	primeMap := make(map[int32]int)
	for i := 'a'; i <= 'z'; i++ {
		primeMap[i] = primeNumber[i-97]
	}
	for _, str := range strs {
		hash := 1
		for i := 0; i < len(str); i++ {
			hash *= primeMap[int32(str[i])]
		}
		resMap[hash] = append(resMap[hash], str)
	}
	for _, v := range resMap {
		res = append(res, v)
	}
	return res
}
