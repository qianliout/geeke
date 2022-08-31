package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{3, 0, 6, 1, 5}
	index := hIndex(citations)
	fmt.Println("index  is ", index)
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
