package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "123"
	target := 6
	res := addOperators2(num, target)
	fmt.Println("res is ", res)
}

/*
282. 给表达式添加运算符
给定一个仅包含数字 0-9 的字符串和一个目标值，在数字之间添加二元运算符（不是一元）+、- 或 * ，返回所有能够得到目标值的表达式。
示例 1:
输入: num = "123", target = 6
输出: ["1+2+3", "1*2*3"]
示例 2:
输入: num = "232", target = 8
输出: ["2*3+2", "2+3*2"]
示例 3:
输入: num = "105", target = 5
输出: ["1*0+5","10-5"]
示例 4:
输入: num = "00", target = 0
输出: ["0+0", "0-0", "0*0"]
示例 5:
输入: num = "3456237490", target = 9191
输出: []
*/
// https://leetcode-cn.com/problems/expression-add-operators/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-52/
func addOperators(num string, target int) []string {
	ans := make([]string, 0)
	addOperatorsHelper([]byte(num), 0, target, "", &ans, 0, 0)
	return ans
}

func addOperatorsHelper(nums []byte, start, target int, path string, res *[]string, pre, value int) {
	if start == len(nums) {
		if value == target {
			*res = append(*res, path)
		}
		return
	}

	for i := start; i < len(nums); i++ {
		// 数字不能从0开始,
		if nums[start] == '0' && i > start {
			break
		}
		cur := genNum(nums, start, i)
		if start == 0 {
			addOperatorsHelper(nums, i+1, target, strconv.Itoa(cur), res, cur, cur)
		} else {
			addOperatorsHelper(nums, i+1, target, path+"+"+strconv.Itoa(cur), res, cur, value+cur)
			addOperatorsHelper(nums, i+1, target, path+"-"+strconv.Itoa(cur), res, -cur, value-cur)
			addOperatorsHelper(nums, i+1, target, path+"*"+strconv.Itoa(cur), res, pre*cur, value-pre+pre*cur)
		}
	}
}

func genNum(nums []byte, start, i int) int {
	v, err := strconv.Atoi(string(nums[start : i+1]))
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	return v
}

func addHelper(nums []byte, start, end, target int, res *[]string) map[string]int {
	opMap := make(map[string]int)

	for i := start + 1; i < end; i++ {
		n, _ := strconv.Atoi(string(nums[start:i]))
		n2, _ := strconv.Atoi(string(nums[i:end]))
		opMap[string(nums[start:i])] = n
		opMap[string(nums[i:end])] = n2

		left := addHelper(nums, start, i, target, res)
		fmt.Println("left ", left)
		right := addHelper(nums, i, end, target, res)

		fmt.Println("right ", right)
		for i, v1 := range left {
			for j, v2 := range right {
				value1 := v1 + v2
				if value1 == target {
					*res = append(*res, i+"+"+j)
				}
				opMap[i+"+"+j] = value1

				value2 := v1 - v2
				if value2 == target {
					*res = append(*res, i+"-"+j)
				}
				opMap[i+"-"+j] = value2
				value3 := v1 * v2
				if value3 == target {
					*res = append(*res, i+"*"+j)
				}
				opMap[i+"*"+j] = value3
			}
		}
	}
	return opMap
}
func addOperators2(num string, target int) []string {
	ans := make([]string, 0)
	addHelper([]byte(num), 0, len(num), target, &ans)
	return ans
}
