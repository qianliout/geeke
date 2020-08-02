package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common"
)

func main() {
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(3, 14))
	fmt.Println(superEggDrop(4, 500))
}

/*
你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
你的目标是确切地知道 F 的值是多少。
无论 F 的初始值如何，你确定 F 的值的最小移动次数是多少？
示例 1：
输入：K = 1, N = 2
输出：2
解释：
鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
如果它没碎，那么我们肯定知道 F = 2 。
因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
示例 2：
输入：K = 2, N = 6
输出：3
示例 3：
输入：K = 3, N = 14
输出：4
提示：
1 <= K <= 100
1 <= N <= 10000
*/
// 这种方法能得到结果,但是会超时
func superEggDrop(K int, N int) int {
	// memo[k][n] 当前状态为 k 个鸡蛋，面对 n 层楼 这个状态下最少的扔鸡蛋次数为 m
	memo := make(map[int]map[int]int)
	for i := 0; i <= K; i++ {
		memo[i] = make(map[int]int)
	}
	return dp(K, N, &memo)
}

// 求的是最小移动次数,而不在那一层破
func dp(K, N int, memo *map[int]map[int]int) int {
	if K == 1 {
		// 只有一个鸡蛋,只能线性扫描
		return N
	}
	if N == 0 {
		return 0
	}

	if v, ok := (*memo)[K][N]; ok {
		return v
	}
	res := math.MaxInt64
	for i := 1; i <= N; i++ {
		res = Min(res, Max(dp(K, N-i, memo), dp(K-1, i-1, memo))+1)
	}
	(*memo)[K][N] = res
	return res
}
