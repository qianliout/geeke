package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cd := Constructor([]int{1, 2, 3})
	fmt.Println(cd.Shuffle())
	fmt.Println(cd.Reset())
}

type Solution struct {
	pre   []int
	after []int
	rd    *rand.Rand
}

func Constructor(nums []int) Solution {
	after := make([]int, len(nums))
	copy(after, nums)
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{
		pre:   nums,
		after: after,
		rd:    rd,
	}

}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.pre
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := len(this.after) - 1; i >= 0; i-- {
		idx := this.rd.Intn(i + 1)
		this.after[i], this.after[idx] = this.after[idx], this.after[i]
	}
	return this.after
}
