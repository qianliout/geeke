package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := Constructor([]int{3, 14, 1, 7})
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())

	// r := find([]int{6, 7, 7, 10}, 8)
	// fmt.Println("r is ", r)
}

/*
给定一个正整数数组 w ，其中 w[i] 代表位置 i 的权重，请写一个函数 pickIndex ，
它可以随机地获取位置 i，选取位置 i 的概率与 w[i] 成正比。

说明:

1 <= w.length <= 10000
1 <= w[i] <= 10^5
pickIndex 将被调用不超过 10000 次
示例1:

输入:
["Solution","pickIndex"]
[[[1]],[]]
输出: [null,0]
示例2:

输入:
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出: [null,0,1,1,1,0]
输入语法说明：
输入是两个列表：调用成员函数名和调用的参数。Solution 的构造函数有一个参数，即数组 w。pickIndex 没有参数。
输入参数是一个列表，即使参数为空，也会输入一个 [] 空列表。
*/

type Solution struct {
	rd    *rand.Rand
	total int64
	node  []Node
}

type Node struct {
	start, end int64
}

func Constructor(w []int) Solution {
	var total int64
	node := make([]Node, len(w))
	for i := 0; i < len(w); i++ {
		node[i] = Node{total, total + int64(w[i])}
		total += int64(w[i])
	}
	return Solution{rand.New(rand.NewSource(time.Now().UnixNano())), total, node}
}

func (this *Solution) PickIndex() int {
	idx := this.rd.Int63n(this.total) % this.total
	for i := 0; i < len(this.node); i++ {
		if this.node[i].start <= idx && this.node[i].end > idx {
			return i
		}
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

// 找靠近右边的一个
func find(num []int, value int) int {
	if len(num) == 0 {
		return 0
	}
	length := len(num) - 1
	left, right := 0, length
	for left < right {
		mid := left + (right-left+1)/2
		if num[mid] > value {
			right = mid - 1
		} else {
			left = mid
		}
	}

	if num[left] >= value {
		return left
	}

	return left + 1
}
