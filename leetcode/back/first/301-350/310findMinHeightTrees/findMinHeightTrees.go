package main

func main() {
	edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}, {6, 5}}
	findMinHeightTrees(6, edges)
}

/*
对于一个具有树特征的无向图，我们可选择任何一个节点作为根。图因此可以成为树，在所有可能的树中，
具有最小高度的树被称为最小高度树。给出这样的一个图，写出一个函数找到所有的最小高度树并返回他们的根节点。
格式
该图包含 n 个节点，标记为 0 到 n - 1。给定数字 n 和一个无向边 edges 列表（每一个边都是一对标签）。
你可以假设没有重复的边会出现在 edges 中。由于所有的边都是无向边， [0, 1]和 [1, 0] 是相同的，
因此不会同时出现在 edges 里。
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
通过次数8,193提交次数24,201
*/

// 循环删除入度为1的点
func findMinHeightTrees(n int, edges [][]int) []int {
	res := make([]int, 0)
	if n == 0 {
		return res
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}
	inDegree := make(map[int][]int)

	// 先计算入度表
	for _, e := range edges {
		first := e[0]
		second := e[1]
		inDegree[first] = append(inDegree[first], second)
		inDegree[second] = append(inDegree[second], first)
	}

	for len(inDegree) > 2 {
		queue := make([][]int, 0)
		for k, v := range inDegree {
			if len(v) == 1 {
				queue = append(queue, []int{k, v[0]})
			}
		}

		// 之后再删除indree,不能再循环中删除(这里是容易出错的地方,不能在循环中删除会导致循环出错)
		for _, q := range queue {
			// 注意这里的理解,这里的first,second和上面是不同的
			first := q[0]
			second := q[1]

			//先删除fist(也就是key)
			delete(inDegree, first)
			// 再删除second中对应是first
			gree := inDegree[second]
			for i, v := range gree {
				if v == first {
					inDegree[second] = append(gree[:i], gree[i+1:]...)
				}
			}
		}
	}
	// 输出结果
	//fmt.Println(inDegree)
	for k, _ := range inDegree {
		res = append(res, k)
	}

	return res
}
