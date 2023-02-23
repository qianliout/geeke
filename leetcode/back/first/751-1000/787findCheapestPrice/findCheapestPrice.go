package main

import (
	"fmt"
)

func main() {
	fligs := [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2}, {4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}
	fmt.Println(findCheapestPrice(5, fligs, 0, 4, 1))
}

/*
有 n 个城市通过 m 个航班连接。每个航班都从城市 u 开始，以价格 w 抵达 v。
现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到从 src 到 dst 最多经过 k 站中转的最便宜的价格。 如果没有这样的路线，则输出 -1。
示例 1：
输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释:
城市航班图如下
示例 2：
输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释:
城市航班图如下
从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
提示：
n 范围是 [1, 100]，城市标签从 0 到 n - 1
航班数量范围是 [0, n * (n - 1) / 2]
每个航班的格式 (src, dst, price)
每个航班的价格范围是 [1, 10000]
k 范围是 [0, n - 1]
航班没有重复，且不存在自环
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	//  dist保存src到各个节点的花费
	dist := make([]int, n)
	inf := 2 * n * 10000
	for i := range dist {
		dist[i] = inf
	}

	for _, f := range flights {
		if f[0] == src {
			dist[f[1]] = f[2]
		}
	}
	for i := 0; i < K; i++ {
		dp := make([]int, n)
		// 这里为什么在用到copy呢
		copy(dp, dist)
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			// 为什么要在上面赋一个最大值呢，比如从0到4，需要经过0->3->4,0是可以到3的，但是
			// dist[u]+w < dp[v] 这上步会报错，因为dp[3]等于0，所以这里应该赋一下最大值
			if dist[u] != inf && dist[u]+w < dp[v] {
				dp[v] = dist[u] + w
			}
		}
		dist = dp
	}
	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}
