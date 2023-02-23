package main

import (
	listnode2 "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

func mergeKLists(lists []*listnode2.ListNode) *listnode2.ListNode {
	if len(lists) == 0 {
		return nil
	}

	root := lists[0]
	for i := 1; i < len(lists); i++ {
		root = mergeTwoLists(root, lists[i])
	}

	return root
}

func mergeTwoLists(l1 *listnode2.ListNode, l2 *listnode2.ListNode) *listnode2.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
