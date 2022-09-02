package main

import (
	"fmt"
	"strconv"
)

func main() {
	ans := addOperators("121", 13)
	fmt.Println("ans is ", ans)
}

func addOperators(num string, target int) []string {
	ans := make([]string, 0)

	dfs([]byte(num), "", &ans, 0, 0, 0, target)
	return ans
}

func dfs(nums []byte, path string, res *[]string, start, pre, cur, target int) {
	if start == len(nums) {
		if cur == target {
			*res = append(*res, path)
		}
		return
	}
	for i := start; i < len(nums); i++ {
		// 不能有前导0
		if nums[start] == '0' && i > start {
			break
		}
		num, _ := strconv.Atoi(string(nums[start : i+1]))
		if start == 0 {
			dfs(nums, string(nums[start:i+1]), res, i+1, num, num, target)
		} else {
			dfs(nums, fmt.Sprintf("%s+%d", path, num), res, i+1, num, cur+num, target)
			dfs(nums, fmt.Sprintf("%s-%d", path, num), res, i+1, -num, cur-num, target)
			dfs(nums, fmt.Sprintf("%s*%d", path, num), res, i+1, pre*num, cur-pre+pre*num, target)
		}
	}
}
