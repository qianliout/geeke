package main

import (
	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {

}

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:
输入: 1->1->1->2->3
输出: 2->3
*/
func deleteDuplicates(head *listnode2.ListNode) *listnode2.ListNode {
	exitMap := make(map[int]int)
	dump := new(listnode2.ListNode)
	dumpCruu := dump
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val && exitMap[cur.Val] == 0 {
			dumpCruu.Next = cur
			exitMap[cur.Val]++
			dumpCruu = dumpCruu.Next
		}
		cur = cur.Next
	}
	dumpCruu.Next = nil
	return dump.Next
}
