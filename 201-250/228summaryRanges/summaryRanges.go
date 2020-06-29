package main

import (
	"fmt"
	"strconv"
)

func main() {

}

/*
给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。
示例 1:
输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。
示例 2:
输入: [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。
*/
func summaryRanges(nums []int) []string {
	res := make([]string, 0)

	length := len(nums)
	if length == 0 {
		return res
	}
	if length == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	start, end := 0, 1
	for end < length {
		if nums[end]-nums[end-1] == 1 {
			end++
		} else {
			if end-1 == start {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
			}

			start = end
			end = start + 1
		}
	}
	//  输出最后一个数
	if end-1 == start {
		res = append(res, strconv.Itoa(nums[start]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
	}

	return res
}

func summaryRanges2(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		res = append(res, strconv.Itoa(nums[0]))
		return res
	}
	start, end := 0, 1
	for end < len(nums) {
		if nums[end]-nums[end-1] == 1 {
			end++
		} else {
			// 只有一个数的情况
			if end-start == 1 {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
			}
			start = end
			end = start + 1
		}
	}
	// 最后一段数据一直是连续的或者最后还有一个数字
	if end-start == 1 {
		res = append(res, strconv.Itoa(nums[start]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
	}

	return res
}
