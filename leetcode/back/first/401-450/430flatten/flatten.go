package main

func main() {
	root := &Node{
		Val:  1,
		Prev: nil,
	}
	next1 := &Node{
		Val:  2,
		Prev: root,
	}
	root.Next = next1
	flatten(root)
}

/*
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。
*/
// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// 其实就是一次遍历，先遍历val,再遍历child，再遍历next,
// 方法一，用slice
func flatten(root *Node) *Node {
	num := traverse(root)
	node := construction(num)
	return node

}

// 先遍历成数组
func traverse(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	child := traverse(root.Child)
	res = append(res, child...)
	next := traverse(root.Next)
	res = append(res, next...)
	return res
}

// 用数组转成双向链表
// 把数组转为双向链表的考法
func construction(num []int) *Node {

	dump := new(Node)
	if len(num) == 0 {
		return dump.Next
	}
	pre := dump
	firstNode := &Node{Val: num[0], Prev: pre}
	pre.Next = firstNode

	for i := 1; i < len(num); i++ {
		nextNode := &Node{Val: num[i]}

		firstNode.Next = nextNode
		nextNode.Prev = firstNode

		pre = firstNode
		firstNode = nextNode
	}

	res := dump.Next
	res.Prev = nil
	return res
}
