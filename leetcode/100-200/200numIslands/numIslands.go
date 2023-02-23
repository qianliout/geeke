package main

func main() {

}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				full(&grid, i, j)
			}
		}
	}
	return ans
}

func full(grid *[][]byte, col, row int) {
	if !in(grid, col, row) {
		return
	}
	if (*grid)[col][row] == '0' {
		return
	}
	(*grid)[col][row] = '0'
	full(grid, col+1, row)
	full(grid, col-1, row)
	full(grid, col, row+1)
	full(grid, col, row-1)
}

func in(grid *[][]byte, col, row int) bool {
	if col < 0 || row < 0 || col > len(*grid)-1 || row > len((*grid)[col])-1 {
		return false
	}
	return true
}
