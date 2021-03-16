package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	blacklist []int
	rd        *rand.Rand
	n         int
}

func Constructor(N int, blacklist []int) Solution {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Solution{
		blacklist: blacklist,
		rd:        rd,
		n:         N,
	}
}

func (this *Solution) Pick() int {
	return 0
}
