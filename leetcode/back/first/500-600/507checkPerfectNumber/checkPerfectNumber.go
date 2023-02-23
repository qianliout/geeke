package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkPerfectNumber(6))
}

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	start := 2
	for start <= int(math.Sqrt(float64(num))) {
		if num%start == 0 {
			sum += start
			sum += num / start
		}
		start++
	}
	if start*start == num {
		sum += start
	}

	return sum == num
}
