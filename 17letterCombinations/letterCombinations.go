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
	res := combinationsUsequeue("23")
	fmt.Println(res)
}

// 回溯的方法，这也是最基本的方法
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

// 队列的方法
// 我们也可以使用队列，先将输入的 digits 中第一个数字对应的每一个字母入队，
// 然后将出队的元素与第二个数字对应的每一个字母组合后入队...直到遍历到 digits
// 的结尾。最后队列中的元素就是所求结果。

func combinationsUsequeue(digits string) []string {
	res := make([]string, 0)
	// var queue [][]byte
	queue := make([][]byte, 1)
	queue[0] = []byte{}
	// queue = append(queue,[][]byte{{''}})

	// queue = append(queue, make([][]byte))

	digitsMap := make(map[byte][]byte)
	digitsMap['2'] = []byte("abc")
	digitsMap['3'] = []byte("def")
	digitsMap['4'] = []byte("ghi")
	digitsMap['5'] = []byte("jkl")
	digitsMap['6'] = []byte("mno")
	digitsMap['7'] = []byte("pqrs")
	digitsMap['8'] = []byte("tuv")
	digitsMap['9'] = []byte("wxyz")

	for i := 0; i < len(digits); i++ {
		for _ = range queue {
			tem := queue[0]
			queue = queue[1:len(queue)]
			liter := digitsMap[digits[i]]
			for _, letter := range liter {
				second := append(tem, letter)
				queue = append(queue, append([]byte{}, second...))
			}

		}

	}

	for _, v := range queue {
		res = append(res, string(v))
	}
	return res
}

/*
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits: return []
        phone = ['abc','def','ghi','jkl','mno','pqrs','tuv','wxyz']
        queue = ['']  # 初始化队列
        for digit in digits:
            for _ in range(len(queue)):
                tmp = queue.pop(0)
                for letter in phone[ord(digit)-50]:# 这里我们不使用 int() 转换字符串，使用ASCII码
                    queue.append(tmp + letter)
        return queue
*/
