package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/
func uniquePaths(m int, n int) int {
	dp := make(map[int]map[int]int)
	// 初始化
	for i := 0; i < m; i++ {
		dp[i] = make(map[int]int)
		dp[i][0] = 1
	}
	// 初值
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]

		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make(map[int]map[int]int)
	// 初始化
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make(map[int]int)
	}

	// 这里的初值一定要写
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = 1
		} else {
			break
		}
	}

	// 这里从1开始，不然就出错
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
