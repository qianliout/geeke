package main

import (
	"fmt"
)

func main() {
	res := summaryRanges([]int{1, 3, 4, 5, 7, 9, 10, 11})
	fmt.Println(res)
}

/*
给定一个  无重复元素 的 有序 整数数组 nums 。
返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，
并且不存在属于某个范围但不属于 nums 的数字 x 。
列表中的每个区间范围 [a,b] 应该按如下格式输出：
"a->b" ，如果 a != b
"a" ，如果 a == b
*/
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	start, end := 0, 0
	for start <= end && end <= len(nums)-1 {
		if end+1 < len(nums) && nums[end+1] == nums[end]+1 {
			end++
		} else {
			if start != end {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end]))
			} else {
				res = append(res, fmt.Sprintf("%d", nums[start]))
			}
			start, end = end+1, end+1
		}
	}
	return res
}
