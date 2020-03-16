package main

import (
	"fmt"
)

func main() {
	x := -121
	res := isPalindrome(x)
	fmt.Println("res is ", res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	res := make([]int, 0)
	for x != 0 {
		pop := x % 10
		res = append(res, pop)
		x = x / 10
	}
	for i := 0; i < len(res); i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}
