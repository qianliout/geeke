package main

import (
	listnode2 "qianliout/leetcode/common/listnode"
)

func main() {

}

/*
编写一个程序，找到两个单链表相交的起始节点。
*/
// 因为题目中说了,一定相交
func getIntersectionNode(headA, headB *listnode2.ListNode) *listnode2.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB
	inversion := 0 // 这里是可以防止没有相交的情况,题目中说了,一定相交,所以,也可以不用这个判断

	for inversion < 3 {
		if currB != currA {
			currB = currB.Next
			if currB == nil {
				currB = headA
				inversion += 1
			}
			currA = currA.Next
			if currA == nil {
				currA = headB
				inversion += 1
			}
		} else {
			return currA
		}
	}
	return nil
}
