package main

import (
	unionfind2 "qianliout/leetcode/leetcode/back/learning/union-find"
)

/*
	There are N students in a class. Some of them are friends, while some are not. Their friendship is

transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C,
then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct
or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1,
then the ith and jth students are direct friends with each other, otherwise not. And you have to output the
total number of friend circles among all the students.

Example 1:

Input:
[[1,1,0],

	[1,1,0],
	[0,0,1]]

Output: 2
Explanation:The 0th and 1st students are direct friends, so they are in a friend circle.
The 2nd student himself is in a friend circle. So return 2.

Example 2:

Input:
[[1,1,0],

	[1,1,1],
	[0,1,1]]

Output: 1
Explanation:The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends,
so the 0th and 2nd students are indirect friends. All of them are in the same friend circle, so return 1.
*/
type Direct [2]int

// findCircleNum  use uniofind method
func findCircleNum(grid [][]int) int {
	uf := unionfind2.NewUnionFind(grid)
	directions := []Direct{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			for _, d := range directions {
				nr, nc := i+d[0], j+d[1]
				if nr >= 0 && nc >= 0 && nr < m && nc < n && grid[nr][nc] == 1 {
					uf.Union(i*n+j, nr*n+nc)
				}
			}
		}
	}

	return uf.Count
}

func FindCircleNumByFull(grid [][]int) int {
	// isViseted := make(map[string]bool)
	count := 0
	for i, v := range grid {
		for j := range v {
			// visKey := fmt.Sprintf("%v_%v", i, j)
			if grid[i][j] == 1 {
				count++
				fullBFS(grid, i, j)
			}
		}
	}
	return count
}

func fullBFS(grid [][]int, i, j int) {
	reFull := make([][2]int, 0)
	reFull = append(reFull, [2]int{i, j})
	for len(reFull) > 0 {
		lenN := len(reFull)
		for k := 0; k < lenN; k++ {
			m := reFull[k][0]
			n := reFull[k][1]
			grid[m][n] = 0

			if m+1 > 0 && m+1 < len(grid) && n > 0 && n < len(grid[0]) && grid[m+1][n] != 0 {
				reFull = append(reFull, [2]int{m + 1, n})
			}
			if m-1 > 0 && m-1 < len(grid) && n > 0 && n < len(grid[0]) && grid[m-1][n] != 0 {
				reFull = append(reFull, [2]int{m - 1, n})
			}
			if m > 0 && m < len(grid) && n+1 > 0 && n+1 < len(grid[0]) && grid[m][n+1] != 0 {
				reFull = append(reFull, [2]int{m, n + 1})
			}
			if m > 0 && m < len(grid) && n-1 > 0 && n-1 < len(grid[0]) && grid[m][n-1] != 0 {
				reFull = append(reFull, [2]int{m, n - 1})
			}
		}
		reFull = reFull[lenN:]
	}
}

// func main() {
//	grid := [][]int{
//		[]int{1, 0, 0, 0, 0},
//		[]int{1, 1, 0, 0, 0},
//		[]int{0, 1, 0, 0, 0},
//		[]int{0, 0, 0, 1, 0},
//	}
//	fmt.Println(FindCircleNumByFull(grid))
// }
