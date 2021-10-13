package main

//
// import (
// 	"fmt"
// 	unionfind "liuqianli/leetcode/union-find"
// )
//
// /*
// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water
// and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are
// all surrounded by water.
// Example 1:
//
// Input:
// 11110
// 11010
// 11000
// 00000
//
// Output: 1
//
// Example 2:
//
// Input:
// 11000
// 11000
// 00100
// 00011
// Output: 3
//
// */
// type Direct [2]int
//
// func numIslands(grid [][]int) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}
//
// 	uf := unionfind.NewUnionFind(grid)
// 	directions := []Direct{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
// 	m, n := len(grid), len(grid[0])
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if grid[i][j] == 0 {
// 				continue
// 			}
// 			for _, d := range directions {
// 				nr, nc := i+d[0], j+d[1]
// 				if nr >= 0 && nc >= 0 && nr < m && nc < n && grid[nr][nc] == 1 {
// 					uf.Union(i*n+j, nr*n+nc)
// 				}
// 			}
// 		}
// 	}
// 	return uf.Count
// }
//
// func main() {
// 	grid := [][]int{
// 		[]int{1, 1, 0, 0, 0},
// 		[]int{1, 1, 0, 0, 0},
// 		[]int{0, 0, 1, 0, 0},
// 		[]int{0, 0, 0, 1, 1},
// 	}
// 	fmt.Println(numIslands(grid))
// }
