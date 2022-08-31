package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo2(0))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(4))
}

func isPowerOfTwo(n int) bool {
	for n != 0 && n&1 == 0 {
		n = n >> 1
	}
	return n == 1
}

func isPowerOfTwo2(n int) bool {
	return n != 0 && n&(n-1) == 0
}
