package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"

	. "qianliout/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	// PrintListNode(head)
	node := reverseBetween(head, 2, 4)
	PrintListNode(node)

}

func test() {

	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if left == 1 {
		// 开始反转
		first := head
		rightNode := first
		for rightNode != nil && right-left >= 0 {
			rightNode = rightNode.Next
			left++
		}
		node := reverse(first, rightNode)
		// 拼接
		cur := node
		for cur != nil && cur.Next != nil && cur != rightNode && cur.Next != rightNode {
			cur = cur.Next
		}
		cur.Next = rightNode
		return node
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverse(head *ListNode, right *ListNode) *ListNode {
	if head == right || head.Next == right {
		return head
	}
	next := reverse(head.Next, right)
	head.Next.Next = head
	head.Next = nil
	return next
}
