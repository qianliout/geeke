package main

import (
	"fmt"
	"regexp"
)

func main() {
	url := "https://access.redhat.com/errata/RHSA-2022-6103.html"
	compile := regexp.MustCompile("/errata/(RHSA-[0-9]+[-:][0-9]+)")
	res := compile.FindStringSubmatch(url)

	// res := majorityElement([]int{1, 2})
	fmt.Println("res ", res)
}

func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	cand1, cand2, count1, count2 := nums[0], nums[0], 0, 0
	for _, num := range nums {
		if num == cand1 {
			count1++
			continue
		}
		if num == cand2 {
			count2++
			continue
		}
		if count1 == 0 {
			cand1 = num
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = num
			count2++
			continue
		}
		count1--
		count2--
	}
	// 计数
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == cand1 {
			count1++
		} else if num == cand2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if count2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}
