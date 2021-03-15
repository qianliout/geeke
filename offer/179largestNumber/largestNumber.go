package main

import (
	"sort"
	"strconv"
	"strings"
)

func main() {
	largestNumber([]int{0, 0})
}

func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		ss[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(Item(ss))
	ans := strings.Join(ss, "")
	if strings.HasPrefix(ans, "0") {
		return "0"
	}
	return ans
}

type Item []string

func (it Item) Len() int {
	return len(it)
}

func (it Item) Less(i, j int) bool {
	return it[i]+it[j] > it[j]+it[i]
}

func (it Item) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}
