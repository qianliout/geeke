package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。
*/
func copyRandomList(head *Node) *Node {
	mem := make(map[*Node]*Node)
	return dfs(head, mem)
}

func dfs(head *Node, mem map[*Node]*Node) *Node {
	if head == nil {
		return head
	}
	if n, ok := mem[head]; ok {
		return n
	}
	newNode := &Node{Val: head.Val}
	mem[head] = newNode
	newNode.Next = dfs(head.Next, mem)
	newNode.Random = dfs(head.Random, mem)
	return newNode
}
