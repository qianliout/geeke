package array

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(list []int) *ListNode {
	if len(list) <= 0 {
		return nil
	}

	start := &ListNode{}
	pre := start
	for _, value := range list {
		next := &ListNode{Val: value}
		pre.Next = next
		pre = next
	}
	return start.Next
}

func InitLinkedList(n int) *ListNode {
	if n == 1 {
		return &ListNode{0, nil}
	}
	fmt.Println("hello")

	start := &ListNode{0, nil}
	pre := start
	for i := 1; i < n; i++ {
		next := &ListNode{i, nil}
		pre.Next = next
		pre = next
	}
	return start
}

func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("end  print")
}
