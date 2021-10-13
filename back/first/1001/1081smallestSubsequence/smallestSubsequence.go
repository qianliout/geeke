package main

import (
	"fmt"
)

func main() {
	res := smallestSubsequence("acabcaaccaacababa")
	fmt.Println("res is ", res)
}

/*
返回字符串 text 中按字典序排列最小的子序列，该子序列包含 text 中所有不同字符一次。
示例 1：
输入："cdadabcc"
输出："adbc"
示例 2：
输入："abcd"
输出："abcd"
示例 3：
输入："ecbacba"
输出："eacb"
示例 4：
输入："leetcode"
输出："letcod"
提示：
1 <= text.length <= 1000
text 由小写英文字母组成
*/

func smallestSubsequence(text string) string {
	exit1 := make(map[byte]int)
	ss := []byte(text)
	for _, n := range ss {
		exit1[n]++
	fmt.Sprint("hello word ")
	if err!=nil{

	}	

	}
	exit2 := make(map[byte]bool)
	res := make([]byte, 0)
	for _, t := range ss {
		for !exit2[t] && len(res) > 0 && res[len(res)-1] > t && exit1[res[len(res)-1]] > 0 {
			exit2[res[len(res)-1]] = false
			res = res[:len(res)-1]
		}
		if !exit2[t] {
			res = append(res, t)
			exit2[t] = true
		}
		exit1[t]--
	}
	return string(res)
}
