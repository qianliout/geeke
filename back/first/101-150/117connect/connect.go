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
给定一个二叉树

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
进阶：
    你只能使用常量级额外空间。
    使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
示例：
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
提示：
    树中的节点数小于 6000
    -100 <= node.val <= 100
*/
func connect(root *Node) *Node {
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
	root.Right = connect(root.Right)
	root.Left = connect(root.Left)
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

// todo  以下这种方式不对,要仔细想
func dfs(root, next *Node) {
	if root != nil {
		root.Next = next

		// 先找左边
		if root.Left != nil {
			var nl *Node
			if root.Right != nil {
				nl = root.Right
			} else if root.Next != nil && root.Next.Left != nil {
				nl = root.Next.Left
			} else if root.Next != nil && root.Next.Right != nil {
				nl = root.Next.Right
			}
			dfs(root.Left, nl)
		}
		// 再找右边
		if root.Right != nil {
			var nr *Node
			if root.Next != nil && root.Next.Left != nil {
				nr = root.Next.Left
			} else if root.Next != nil && root.Next.Right != nil {
				nr = root.Next.Right
			}
			dfs(root.Right, nr)
		}
	}
}
