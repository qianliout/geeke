package main

import "fmt"

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例:
输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
func main() {
	res := letterCombinations("23")
	fmt.Println(res)

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	// var digitsByteSlice = make([][]byte, 0)
	digitsMap := make(map[byte][]byte)
	digitsMap['2'] = []byte("abc")
	digitsMap['3'] = []byte("def")
	digitsMap['4'] = []byte("ghi")
	digitsMap['5'] = []byte("jkl")
	digitsMap['6'] = []byte("mno")
	digitsMap['7'] = []byte("pqrs")
	digitsMap['8'] = []byte("tuv")
	digitsMap['9'] = []byte("wxyz")

	// fmt.Println("digitsByteSlice is ", digitsByteSlice)

	path := make([]byte, 0)
	// used := make([][4]bool, 8)
	combinations([]byte(digits), digitsMap, path, 0, &res)
	return res
}

// depth 表示对应的数字
func combinations(digits []byte, digitsMap map[byte][]byte, path []byte, depth int32, res *[]string) {
	if len(path) == len(digits) {
		// fmt.Println("res:", string(path))
		*res = append(*res, string(append([]byte{}, path...)))
		return
	}
	inters := digitsMap[digits[depth]]

	for _, char := range inters {
		path = append(path, char)
		combinations(digits, digitsMap, path, depth+1, res)
		path = path[:len(path)-1]
	}
}
