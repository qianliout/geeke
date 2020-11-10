package main

import (
	"sort"

	. "outback/leetcode/common/listnode"
)

func main() {
	list1 := ListNode{Val: 2}
	//list1.Next = &ListNode{Val: 4}
	//list1.Next.Next = &ListNode{Val: 5}

	list2 := ListNode{Val: -1}
	//list2.Next = &ListNode{Val: 4}
	//list2.Next.Next = &ListNode{Val: 6}
	res := mergeKLists([]*ListNode{&list1, nil, &list2})
	PrintListNode(res)
}

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return Divide(lists, 0, len(lists)-1)
}

// 暴力求解
func Violence(lists []*ListNode) *ListNode {

	root := new(ListNode)
	allValue := make([]int, 0)

	for _, node := range lists {
		for node != nil {
			allValue = append(allValue, node.Val)
			node = node.Next
		}
	}
	sort.Ints(allValue)
	curr := root
	for _, v := range allValue {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return root.Next
}

// 每一个逐一比较，每一次比较找出最小的那个值
func OneByOneToCompare(lists []*ListNode) *ListNode {
	result := new(ListNode)

	for len(lists) > 0 {
		curr := new(ListNode)
		currNode := 0
		for i, node := range lists {
			if node.Next == nil {
				lists = append(append([]*ListNode{}, lists[:i]...), lists[i+1:]...)
			}
			if node.Val < curr.Val {
				curr.Val = node.Val
				currNode = i
			}
		}
		lists[currNode] = lists[currNode].Next
		result.Next = curr
	}
	return result.Next
}

func mergeOneByOne(lists []*ListNode) *ListNode {
	var root *ListNode
	for _, node := range lists {
		if node != nil {
			root = mergeTwoLists(root, node)
		}
	}
	return root
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else {
		if l1.Val <= l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoLists(l2.Next, l1)
			return l2
		}
	}
}

// 分治方法，也是二分查找法
func Divide(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	l1 := Divide(lists, left, mid)
	l2 := Divide(lists, mid+1, right)
	return mergeTwoLists(l1, l2)
}
