package main

import (
	"outback/leetcode/back/common/listnode"
)

func main() {

}

/*
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
现有一个链表 -- head = [4,5,1,9]，它可以表示为:
示例 1:
输入: head = [4,5,1,9], node = 5
输出: [4,1,9]
解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:
输入: head = [4,5,1,9], node = 1
输出: [4,5,9]
解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
*/

/*
这道题是真的精妙，不懂为啥还有那么多吐槽的哈哈！
node这个节点就是需要删除的节点；之前我们可以用head->next->val去判断下一个是否是删除的节点，
然后head->next=head->next->next，这题可以用把 node下一节点复制到node，把下一节点跳过！这道题出的挺好的啊！
*/
func deleteNode(node *listnode.ListNode, value int) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
