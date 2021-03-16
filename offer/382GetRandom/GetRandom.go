package main

import (
	"math/rand"
	"time"

	. "outback/leetcode/common/listnode"
)

func main() {

}

type Solution struct {
	head *ListNode
	rd   *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Solution{
		head: head,
		rd:   rd,
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	count, ans, cur := 0, 0, this.head
	for cur != nil {
		count++
		if count == this.rd.Intn(count)+1 {
			ans = cur.Val
		}
		cur = cur.Next
	}
	return ans
}
