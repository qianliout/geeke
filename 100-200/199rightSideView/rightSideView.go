package main

import (
	"fmt"
	"time"

	. "qianliout/leetcode/common/treenode"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go Test(ch1, ch2)
	go func() {
		ch1 <- 1
	}()

	time.Sleep(time.Minute)

}

func Test(ch1, ch2 chan int) {
	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	}
}

/*
199. 二叉树的右视图
*/
func rightSideView2(root *TreeNode) []int {
	ans, queue := make([]int, 0), make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		tem := make([]*TreeNode, 0)
		for i := range queue {
			node := queue[i]
			if node.Left != nil {
				tem = append(tem, node.Left)
			}
			if node.Right != nil {
				tem = append(tem, node.Right)
			}
		}
		queue = tem
	}
	return ans
}
