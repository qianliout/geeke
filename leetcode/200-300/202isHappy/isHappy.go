package main

import (
	"fmt"
)

func main() {
	happy := isHappy(19)
	fmt.Println(happy)
}

func isHappy(n int) bool {
	exit := make(map[int]struct{})
	for n > 1 {
		n = happy(n)
		if _, ok := exit[n]; ok {
			return false
		}
		exit[n] = struct{}{}
	}
	return true
}

func happy(n int) int {
	ans := make([]int, 0)
	for n > 0 {
		ans = append(ans, n%10)
		n = n / 10
	}

	res := 0
	for i := 0; i < len(ans); i++ {
		res = res + ans[i]*ans[i]
	}
	return res
}
