package main

import (
	"qianliout/leetcode/back/common/treenode"
)

func main() {

}

/*
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，
以便稍后在同一个或另一个计算机环境中重建。
设计一个算法来序列化和反序列化二叉搜索树。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，
并且可以将该字符串反序列化为最初的二叉搜索树。
编码的字符串应尽可能紧凑。
注意：不要使用类成员/全局/静态变量来存储状态。 你的序列化和反序列化算法应该是无状态的。
*/

//  因为二叉搜索树也是二叉对,所以使用297的方法可解
type Codec struct {
}

func Constructor() Codec {

}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *treenode.TreeNode) string {

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *treenode.TreeNode {

}
