package main

import (
	"fmt"
)

func main() {
	res := canIWin(10, 11)
	fmt.Println("can i win is ", res)
}

/*
在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和达到 100 的玩家，即为胜者。
如果我们将游戏规则改为 “玩家不能重复使用整数” 呢？
例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
给定一个整数 maxChoosableInteger （整数池中可选择的最大数）和另一个整数 desiredTotal（累计和），
判断先出手的玩家是否能稳赢（假设两位玩家游戏时都表现最佳）？
你可以假设 maxChoosableInteger 不会大于 20， desiredTotal 不会大于 300。
示例：
输入：
maxChoosableInteger = 10
desiredTotal = 11
输出：
false
解释：
无论第一个玩家选择哪个整数，他都会失败。
第一个玩家可以选择从 1 到 10 的整数。
如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。
*/

//dp
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*((maxChoosableInteger+1)/2) < desiredTotal {
		return false
	}
	if maxChoosableInteger >= desiredTotal {
		return true
	}

	dp := make([]int, 1<<uint(maxChoosableInteger)-1)
	return dfs(dp, maxChoosableInteger, 0, desiredTotal)
}

func dfs(dp []int, maxChoosableInteger, v, desiredTotal int) bool {
	if dp[v] != 0 {
		return dp[v] == 1
	}
	dp[v] = 2
	for i := 1; i <= maxChoosableInteger; i++ {
		selectNum := 1<<uint(i) - 1
		if selectNum&v > 0 {
			continue
		}
		if i >= desiredTotal {
			dp[v] = 1
			break
		}
		next := dfs(dp, maxChoosableInteger, v|selectNum, desiredTotal-i)
		if !next {
			dp[v] = 1
			break
		}
	}
	return dp[v] == 1
}
