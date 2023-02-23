package main

import (
	"fmt"

	listnode2 "qianliout/leetcode/leetcode/common/listnode"
)

func main() {
	fmt.Println(splitListToParts())
}

/*
给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。
每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。
这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。
返回一个符合上述规则的链表的列表。
举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]
示例 1：
输入:
root = [1, 2, 3], k = 5
输出: [[1],[2],[3],[],[]]
解释:
输入输出各部分都应该是链表，而不是数组。
例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且 root.next.next.next = null。
第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。
示例 2：
输入:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
解释:
输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。
提示:
root 的长度范围： [0, 1000].
输入的每个节点的大小范围：[0, 999].
k 的取值范围： [1, 50].
*/

// 可以获取正确的解，但是解法太不优雅了
func splitListToParts(root *listnode2.ListNode, k int) []*listnode2.ListNode {
	// 计算原链表的长度
	count := 0
	nod := root
	for nod != nil {
		count++
		nod = nod.Next
	}
	var inter int
	if count%k == 0 {
		inter = count / k
	} else {
		inter = count/k + 1
	}
	res := make([]*listnode2.ListNode, 0)
	nod = root
	start := 0
	prek := k
	startNode := nod
	for nod != nil {
		start++
		if start == inter {
			nex := nod.Next
			nod.Next = nil
			nod = nex
			res = append(res, startNode)
			startNode = nod
			start = 0
			k = k - 1
			if k == 0 {
				break
			}
			count -= inter
			if count%k == 0 {
				inter = count / k
			} else {
				inter = count/k + 1
			}
			continue
		}
		nod = nod.Next
	}
	res = append(res, startNode)
	if len(res) > prek {
		res = res[:prek]
	}

	for len(res) < prek {
		res = append(res, nil)
	}

	return res
}

func splitListToParts2(root *listnode2.ListNode, k int) []*listnode2.ListNode {
	count := 0
	tem := root
	for tem != nil {
		tem = tem.Next
		count++
	}
	res := make([]*listnode2.ListNode, k)
	if count <= 0 || k <= 0 {
		return res
	}
	inter := 0
	if count <= k {
		inter = 1
	} else {
		inter = count / k
	}
	// over表示有多个少元素要增加一
	over := 0
	if count > k {
		over = count % k
	}
	for i := 0; i < k; i++ {
		res[i] = root
		// 移动指针
		for j := 0; j < inter-1; j++ {
			if root == nil {
				return res
			}
			root = root.Next
		}
		if over > 0 {
			root = root.Next
			over--
		}
		if root == nil {
			break
		}
		// 断开前保存下一个节点
		next := root.Next
		// 部分之间要断开链表
		root.Next = nil
		root = next
	}
	return res
}
