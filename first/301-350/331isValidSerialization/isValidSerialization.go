package main

import (
	"strings"
)

func main() {

}

/*
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。
如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
示例 1:
输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
示例 2:
输入: "1,#"
输出: false
示例 3:
输入: "9,#,#,1"
输出: false
*/
// 消耗槽位思想
func isValidSerialization(preorder string) bool {
	order := strings.Split(preorder, ",")
	if len(order) == 0 {
		return false
	}

	slots := 1
	for _, o := range order {
		slots--
		if slots < 0 {
			return false
		}
		if o != "#" {
			slots = slots + 2
		}

	}
	return slots == 0
}

var i int

// 模拟建树过程, 这个方法在go语言中没有通过,不知道是为什么
func isValidSerialization2(preorder string) bool {
	order := strings.Split(preorder, ",")
	//构建完整棵树以后，序列化中的所有结点都应被用到，也就是 i == strings.length
	return canBuid(order) && i == len(order)
}

func canBuid(order []string) bool {
	//如果在建立当前结点时发现序列化中没有结点了，那就返回 false
	if i == len(order) {
		return false
	}
	//如果当前结点是 null，就返回 true
	if order[i] == "#" {
		return true
	}
	i++
	//构建左子树和右子树
	return canBuid(order) && canBuid(order)
}
