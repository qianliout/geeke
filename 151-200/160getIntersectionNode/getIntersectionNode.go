package main

import (
	. "outback/leetcode/common/listnode"
)

func main() {

}

// 因为题目中说了,一定相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB
	inversion := 0

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
