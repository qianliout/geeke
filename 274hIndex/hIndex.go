package main

import (
	"fmt"
	"sort"

	"qianliout/leetcode/274hIndex/name"
)

func main() {

	r := &name.Res{}

	r.SetAge(12)
	fmt.Println(r.GetAge())
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := len(citations)
	for i := range citations {
		if citations[i] < h {
			h--
		}
	}
	return h
}
