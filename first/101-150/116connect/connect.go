package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
示例：
输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}
输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
提示：
    你只能使用常量级额外空间。
    使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/
// 因为是完美二叉树,所以可以直接递归
func connect(root *Node) *Node {
	dfs(root, nil)
	return root
}

// 因为是完美二叉树
func dfs(root, next *Node) {
	if root != nil {
		root.Next = next
		dfs(root.Left, root.Right)
		var n *Node
		if root.Next != nil {
			n = root.Next.Left
		}
		dfs(root.Right, n)
	}
}

func connect2(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
		root.Right.Next = getNextNode(root)
	}
	if root.Left == nil && root.Right != nil {
		root.Right.Next = getNextNode(root)
	}
	if root.Right == nil && root.Left != nil {
		root.Left.Next = getNextNode(root)
	}
	// 这里要注意：先递归右子树，否则右子树根节点next关系没建立好，左子树到右子树子节点无法正确挂载
	root.Right = connect2(root.Right)
	root.Left = connect2(root.Left)
	return root
}

func getNextNode(root *Node) *Node {
	for root.Next != nil {
		if root.Next.Left != nil {
			return root.Next.Left
		}
		if root.Next.Right != nil {
			return root.Next.Right
		}
		root = root.Next
	}
	return nil
}
