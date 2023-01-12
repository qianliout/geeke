package main

import (
	"fmt"
	"sort"
)

func main() {
	// hand := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	// hand := []int{1, 2, 3, 4, 5}
	hand := []int{2, 1}

	groupSize := 2
	ans := isNStraightHand(hand, groupSize)
	fmt.Println("ans is ", ans)
}

func isNStraightHand(hand []int, groupSize int) bool {
	sort.Ints(hand)
	left := make(map[int]int)
	for i := range hand {
		left[hand[i]]++
	}
	for _, v := range hand {
		if left[v] <= 0 {
			continue
		}
		// 否则，生成一个长度为 groupSize 的连续子序列
		for i := 1; i < groupSize; i++ {
			if left[v+i] <= 0 {
				return false
			}
		}
		for i := 0; i < groupSize; i++ {
			left[v+i]--
		}
	}
	return true
}
