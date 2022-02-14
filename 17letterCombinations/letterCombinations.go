package main

import (
	"fmt"
)

func main() {
	combinations := userQueue("23")
	fmt.Println(combinations)
}

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
func letterCombinations(digits string) []string {
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	path, res := make([]byte, 0), make([]string, 0)
	// 本地执行正确，测试用例测试出错，会输出 [""]
	if digits == "" {
		return res
	}

	backtrack([]byte(digits), 0, path, digitMap, &res)
	return res
}

func backtrack(nums []byte, start int, path []byte, digitMap map[byte][]byte, res *[]string) {
	if len(path) == len(nums) {
		*res = append(*res, string(append([]byte{}, path...)))
		return
	}
	for _, b := range digitMap[nums[start]] {
		path = append(path, b)
		backtrack(nums, start+1, path, digitMap, res)
		path = path[:len(path)-1]
	}
}

// use queue
func userQueue(digits string) []string {
	res := make([]string, 0)
	if digits == "" {
		return res
	}
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	queue := make([][]byte, 0)
	queue = append(queue, []byte{})
	ss := []byte(digits)
	for i := 0; i < len(ss); i++ {
		length := len(queue)
		for j := 0; j < length; j++ {
			first := queue[0]
			queue = queue[1:]
			letter := digitMap[ss[i]]
			for _, b := range letter {
				queue = append(queue, append([]byte{}, append(first, b)...))
			}
		}
	}

	for _, v := range queue {
		res = append(res, string(v))
	}
	return res
}
