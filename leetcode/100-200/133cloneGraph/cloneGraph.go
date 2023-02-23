package main

func main() {

}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	mem := make(map[*Node]*Node)
	return dfs(node, mem)
}

func dfs(node *Node, mem map[*Node]*Node) *Node {
	if node == nil {
		return node
	}
	if pre, ok := mem[node]; ok {
		return pre
	}
	nowNod := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	mem[node] = nowNod

	for _, n := range node.Neighbors {
		nowNod.Neighbors = append(nowNod.Neighbors, dfs(n, mem))
	}
	return nowNod
}
