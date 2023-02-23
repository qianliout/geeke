package main

import (
	. "qianliout/leetcode/leetcode/common/listnode"
)

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	max, currA, currB := 0, headA, headB

	for max < 3 {
		if currA == currB {
			return currA
		} else if currA == nil {
			currA = headB
			max++
		} else if currB == nil {
			currB = headA
			max++
		} else {
			currB = currB.Next
			currA = currA.Next
		}
	}
	return nil
}
