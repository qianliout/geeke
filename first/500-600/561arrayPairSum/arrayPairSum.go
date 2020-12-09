package main

import (
	"sort"
)

func main() {

}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i = i + 2 {
		ans += nums[i]
	}
	return ans
}
