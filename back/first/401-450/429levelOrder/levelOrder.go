package main

func main() {

}

/*
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
例如，给定一个 3叉树 :
返回其层序遍历:
[
     [1],
     [3,2,4],
     [5,6]
]
说明:
    树的深度不会超过 1000。
    树的节点总数不会超过 5000。
*/
// Definition for a Node.

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		thisRes := make([]int, 0)
		nextLevel := make([]*Node, 0)

		for _, n := range queue {
			thisRes = append(thisRes, n.Val)
			nextLevel = append(nextLevel, n.Children...)
		}

		res = append(res, thisRes)
		queue = nextLevel
	}
	return res
}
