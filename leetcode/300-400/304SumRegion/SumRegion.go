package main

func main() {

}

type NumMatrix struct {
	Dp     [][]int
	Matrix [][]int
	UseDp  bool
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{Matrix: matrix, Dp: matrix, UseDp: false}
	}

	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	// 初值
	dp[0][0] = 0 // 本身就有默认值
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + matrix[i-1][0]
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i-1]
	}

	// 求值
	for col := 1; col < len(dp); col++ {
		for row := 1; row < len(dp[col]); row++ {
			dp[col][row] = dp[col-1][row] + dp[col][row-1] - dp[col-1][row-1] + matrix[col-1][row-1]
		}
	}
	return NumMatrix{Matrix: matrix, Dp: dp, UseDp: true}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if !this.UseDp {
		return 0
	}
	return this.Dp[row2+1][col2+1] + this.Dp[row1][col1] - this.Dp[row2+1][col1] - this.Dp[row1][col2+1]
}
