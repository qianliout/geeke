package main

import (
	"fmt"
	"math"
)

func main() {
	res := findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})
	fmt.Println("res is ", res)
}

/*
对于一个具有树特征的无向图，我们可选择任何一个节点作为根。图因此可以成为树，在所有可能的树中，具有最小高度的树被称为最小高度树。给出这样的一个图，
写出一个函数找到所有的最小高度树并返回他们的根节点。
格式
该图包含 n 个节点，标记为 0 到 n - 1。给定数字 n 和一个无向边 edges 列表（每一个边都是一对标签）。
你可以假设没有重复的边会出现在 edges 中。由于所有的边都是无向边， [0, 1]和 [1, 0] 是相同的，因此不会同时出现在 edges 里。
示例 1:
输入: n = 4, edges = [[1, 0], [1, 2], [1, 3]]
        0
        |
        1
       / \
      2   3
输出: [1]
示例 2:
输入: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]
     0  1  2
      \ | /
        3
        |
        4
        |
        5
输出: [3, 4]
说明:
     根据树的定义，树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
    树的高度是指根节点和叶子节点之间最长向下路径上边的数量。
*/
func findMinHeightTrees(n int, edges [][]int) []int {
	inDegree := make(map[int][]int)
	for i := 0; i < n; i++ {
		inDegree[i] = make([]int, 0)
	}
	// 计算入度
	for _, e := range edges {
		inDegree[e[1]] = append(inDegree[e[1]], e[0])
		inDegree[e[0]] = append(inDegree[e[0]], e[1])
	}
	for len(inDegree) > 2 {
		fmt.Println("hello ", inDegree)
		queue := make([][]int, 0)
		for k, v := range inDegree {
			if len(v) == 1 {
				queue = append(queue, []int{k, v[0]})
			}
		}
		for _, q := range queue {
			first := q[0]
			second := q[1]

			delete(inDegree, first)

			gre := inDegree[second]
			for i, k := range gre {
				if k == first {
					// 这里使用到深复制
					inDegree[second] = append(gre[:i], gre[i+1:]...)
				}
			}
		}
	}

	res := make([]int, 0)
	for k, _ := range inDegree {
		res = append(res, k)
	}
	return res
}

// 如果是返回任意一个的话，这种方法就可以
func findMinHeightTrees2(n int, edges [][]int) []int {
	ingree := make(map[int][]int)
	for _, n := range edges {
		ingree[n[0]] = append(ingree[n[0]], n[1])
		ingree[n[1]] = append(ingree[n[1]], n[0])
	}
	// 再去个重排个序不就得了吗
	number := make(map[int]int)
	max := math.MinInt64
	for n, m := range ingree {
		ans := dup(m)
		ingree[n] = ans
		number[n] = len(ans)
		if len(ans) > max {
			max = len(ans)
		}
	}
	ans := make([]int, 0)
	for k, v := range number {
		if v == max {
			ans = append(ans, k)
		}
	}
	return ans
}

func dup(pre []int) []int {
	ans := make([]int, 0)
	m := make(map[int]bool)
	for _, n := range pre {
		if !m[n] {
			ans = append(ans, n)
			m[n] = true
		}
	}
	return ans
}
