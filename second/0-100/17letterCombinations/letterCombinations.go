package main

import (
	"fmt"
)

func main() {
	res := usequeu("23")
	fmt.Println("res is ", res)
}

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

func letterCombinations(digits string) []string {
	digitsMap := make(map[byte][]byte)
	digitsMap['2'] = []byte("abc")
	digitsMap['3'] = []byte("def")
	digitsMap['4'] = []byte("ghi")
	digitsMap['5'] = []byte("jkl")
	digitsMap['6'] = []byte("mno")
	digitsMap['7'] = []byte("pqrs")
	digitsMap['8'] = []byte("tuv")
	digitsMap['9'] = []byte("wxyz")

	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	used := make(map[int]map[byte]bool)
	for i := 0; i < len(digits); i++ {
		used[i] = map[byte]bool{}
	}

	dfs(digits, 0, digitsMap, &used, []byte{}, &res)
	return res
}

func dfs(digits string, n int, digitsMap map[byte][]byte, used *map[int]map[byte]bool, path []byte, res *[]string) {
	if len(path) == len(digits) && n >= len(digits) {
		*res = append(*res, string(path))
		return
	}
	for _, c := range digitsMap[digits[n]] {
		if !(*used)[n][c] {
			path = append(path, c)
			(*used)[n][c] = true
			dfs(digits, n+1, digitsMap, used, path, res)
			path = path[:len(path)-1]
			(*used)[n][c] = false
		}
	}
}

// 使用队列
// 队列的方法
// 我们也可以使用队列，先将输入的 digits 中第一个数字对应的每一个字母入队，
// 然后将出队的元素与第二个数字对应的每一个字母组合后入队...直到遍历到 digits
// 的结尾。最后队列中的元素就是所求结果。
func usequeu(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	digitsMap := make(map[byte][]byte)
	digitsMap['2'] = []byte("abc")
	digitsMap['3'] = []byte("def")
	digitsMap['4'] = []byte("ghi")
	digitsMap['5'] = []byte("jkl")
	digitsMap['6'] = []byte("mno")
	digitsMap['7'] = []byte("pqrs")
	digitsMap['8'] = []byte("tuv")
	digitsMap['9'] = []byte("wxyz")

	queue := make([][]byte, 1)
	queue[0] = []byte{}
	for i := 0; i < len(digits); i++ {
		second := make([][]byte, 0)
		leght := len(queue)
		for j := 0; j < leght; j++ {
			q := queue[0]
			queue = queue[1:]
			for _, t := range digitsMap[digits[i]] {
				s := append(q, t)
				// 下面这一步是最容易出错的,因为，如果q,没有扩容的话，s用的是q的地址
				second = append(second, append([]byte{}, s...))
			}
		}
		queue = second
	}

	for _, q := range queue {
		res = append(res, string(q))
	}
	return res
}
