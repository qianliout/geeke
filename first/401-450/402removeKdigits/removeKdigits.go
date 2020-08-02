package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := "1432219"
	kdigits := removeKdigits2(nums, 3)

	fmt.Println("res is ", kdigits)
}

/*
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:
    num 的长度小于 10002 且 ≥ k。
    num 不会包含任何前导零。
示例 1 :
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/

func removeKdigits(num string, k int) string {
	nums := []byte(num)
	number := make([]int, 0)
	for _, a := range nums {
		i, _ := strconv.Atoi(string(a))
		number = append(number, i)
	}
	deleteIndex := make(map[int]bool, 0)
	stark := make([]int, 0)
	for i, n := range number {
		if k == 0 {
			break
		}
		if len(stark) <= 0 {
			stark = append(stark, i)
		} else {
			for len(stark) > 0 {
				last := stark[len(stark)-1]
				if number[last] > n {
					deleteIndex[last] = true
					stark = stark[:len(stark)-1]
					k--
					if k == 0 {
						break
					}
				} else {
					break
				}
			}
			stark = append(stark, i)
		}
	}
	//fmt.Println("deleteIndex", deleteIndex)
	//fmt.Println("stark", stark, k)
	// 循环结束,k 不为0,且stark不为空
	for k > 0 && len(stark) > 0 {
		deleteIndex[stark[len(stark)-1]] = true
		stark = stark[:len(stark)-1]
		k--
	}
	//fmt.Println("deleteIndex", deleteIndex)
	//fmt.Println("stark", stark, k)

	// 输出结果
	first := false
	ans := make([]int, 0)
	for i := range number {
		if !deleteIndex[i] {
			if number[i] == 0 && !first {
				continue
			}
			first = true
			ans = append(ans, number[i])
		}
	}
	//fmt.Println("ans", ans)

	// 把前面的所有0去了
	s := ""
	for _, a := range ans {
		s += strconv.Itoa(a)
	}
	if len(s) == 0 {
		return "0"
	}
	return s
}

func removeKdigits2(num string, k int) string {

	nums := make([]int, 0)
	remind := len(num) - k
	for i := range num {
		n, _ := strconv.Atoi(string(num[i]))
		nums = append(nums, n)
	}

	stack := make([]int, 0)
	for _, n := range nums {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > n {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, n)
	}
	// 输出结果
	res := ""
	start := true
	for j := 0; j < remind; j++ {
		if start && stack[j] == 0 {
			continue
		}
		start = false
		res = res + strconv.Itoa(stack[j])
	}
	if len(res) <= 0 {
		return "0"
	}

	return res
}
