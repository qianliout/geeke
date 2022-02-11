package main

import (
	"fmt"
	"math/rand"
	"time"

	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {
	head := &listnode2.ListNode{Val: 10}
	head.Next = &listnode2.ListNode{Val: 100}
	head.Next.Next = &listnode2.ListNode{Val: 100}
	head.Next.Next.Next = &listnode2.ListNode{Val: 20}
	head.Next.Next.Next.Next = &listnode2.ListNode{Val: 20}
	head.Next.Next.Next.Next.Next = &listnode2.ListNode{Val: 100}
	s := Constructor(head)
	ran := s.GetRandom()
	fmt.Println("rand is ", ran)
	fmt.Println(rand.Intn(1))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。
进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？
示例:
// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);
// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();
*/
type Solution struct {
	head *listnode2.ListNode

	r *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *listnode2.ListNode) Solution {
	seed := time.Now().UnixNano()
	ran := rand.New(rand.NewSource(seed))
	return Solution{head: head, r: ran}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	count := 1
	cur := this.head
	reverse := cur.Val
	for cur != nil {
		if count == this.r.Intn(count)+1 {
			reverse = cur.Val
		}
		count++
		cur = cur.Next
	}
	return reverse
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
