package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3, 1, 2, 3}
	nums := []int{1, 2, 3, 1}
	duplicate := containsNearbyDuplicate(nums, 3)
	fmt.Println("dup", duplicate)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	exit := make(map[int]int)
	for i, v := range nums {
		if pre, ok := exit[v]; ok && i-pre <= k {
			return true
		}
		exit[v] = i
	}
	return false
}
