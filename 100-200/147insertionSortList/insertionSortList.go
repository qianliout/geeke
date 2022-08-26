package main

import (
	. "qianliout/leetcode/common/listnode"
)

func main() {
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	node := insertionSortList(head)
	PrintListNode(node)

}

// 对链表进行插入排序,升序
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dump := new(ListNode)
	for head != nil {
		cur := head
		head = head.Next // 只能入在这里，放在最下边就错，要注意理解。因为cur和head共用一是一个节点地址，改变了cur.Next就改变了head
		node := search(dump, cur)
		cur.Next = node.Next
		node.Next = cur
	}
	return dump.Next
}

// 对链表进行插入排序,升序
func insertionSortList2(head *ListNode) *ListNode {
	return dfs(head)
}

// root是已排序的链表，所以找到第一个比node大的链表的节点，返回这个节点的前一个节点
func search(root *ListNode, node *ListNode) *ListNode {
	for root != nil && root.Next != nil {
		if root.Next.Val >= node.Val {
			break
		}
		root = root.Next
	}
	return root
}

func dfs(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}

	// 这种方法一定要注意
	dump := new(ListNode)

	next := dfs(root.Next)
	dump.Next = next
	// 找到合适的地方，放好
	pre, greater := dump, next
	for greater != nil && root.Val > greater.Val {
		pre = pre.Next
		greater = greater.Next
	}
	pre.Next = root
	root.Next = greater
	return dump.Next
}
