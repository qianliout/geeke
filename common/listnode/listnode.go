package listnode

import (
	"fmt"
	"strconv"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

type DoubleLinkedNode struct {
	Key  int
	Val  int
	Pre  *DoubleLinkedNode
	Post *DoubleLinkedNode
}

func NewDumpDoubleLinkedNode() *DoubleLinkedNode {
	return new(DoubleLinkedNode)
}

func NewDoubleLinkedNode(key, value int) *DoubleLinkedNode {
	return &DoubleLinkedNode{
		Key: key,
		Val: value,
	}
}

// 添加节点到开头,并返回头节点
func (d *DoubleLinkedNode) AddFirst(n *DoubleLinkedNode) *DoubleLinkedNode {
	d.Print("AddFirst befor " + strconv.Itoa(n.Key) + " " + strconv.Itoa(n.Val))
	if d == nil {
		d = n
		return n
	}
	d.Pre = n
	n.Post = d
	d = n
	// d.Print("AddFirst after " + strconv.Itoa(n.Key) + " " + strconv.Itoa(n.Val))
	return d
}

// 移出一个节点
func (d *DoubleLinkedNode) Remove(n *DoubleLinkedNode) *DoubleLinkedNode {
	// d.Print("befor remove:" + strconv.Itoa(n.Key) + " " + strconv.Itoa(n.Val))
	if n == nil {
		return d
	}
	if n.Pre == nil {
		d = n.Post
		if d != nil {
			d.Pre = nil
		}
		// d.Print("after remove:")
		return d
	}
	if n.Post == nil {
		n.Pre.Post = nil
		// d.Print("after remove:")
		return d
	}
	n.Post.Pre = n.Pre
	n.Pre.Post = n.Post

	// d.Print("after remove:")
	return d
}

// 删除最后一个节点,并返回该节点
func (d *DoubleLinkedNode) PopLast() *DoubleLinkedNode {
	cur := d
	for cur != nil && cur.Post != nil {
		cur = cur.Post
	}
	if cur == nil || cur.Pre == nil {
		return nil
	}
	cur.Pre.Post = nil
	return cur
}

func (d *DoubleLinkedNode) Size() int {
	return 0
}

func (d *DoubleLinkedNode) Print(str string) {
	cur := d
	res := make([][]int, 0)
	for cur != nil {
		res = append(res, []int{cur.Key, cur.Val})
		cur = cur.Post
	}
	fmt.Println(str, "double:", res)
}
