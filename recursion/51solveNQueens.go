package recursion

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return [][]string{}
	}
	var re [][]int

	//三个map，shus就是竖，扑面而来的爱国情怀。

	shus, pies, nas := make(map[int]bool, n), make(map[int]bool, n), make(map[int]bool, n)
	DFS := func(rows []int, n int) {}
	DFS = func(rows []int, n int) {
		row := len(rows)
		if row == n {
			aaaa := make([]int, len(rows))
			copy(aaaa, rows)
			//re = append(re,append([]int{},rows...))
			re = append(re, aaaa)
			return
		}

		for col := 0; col < n; col++ {
			if !shus[col] && !pies[row+col-1] && !nas[row-col-1] {
				shus[col] = true
				pies[row+col-1] = true
				nas[row-col-1] = true
				DFS(append(rows, col), n)
				shus[col] = false
				pies[row+col-1] = false
				nas[row-col-1] = false
			}
		}
	}

	DFS([]int{}, n)
	return bQ(re, n)
}

func bQ(re [][]int, n int) (result [][]string) {
	for _, v := range re {
		s := make([]string, 0)
		for _, vv := range v {
			str := ""
			for i := 0; i < n; i++ {
				if i == vv {
					str += "Q"
				} else {
					str += "."
				}
			}
			s = append(s, str)
		}
		result = append(result, s)
	}
	return
}

func SolveNQueens(n int) [][]string {
	reslut := make([][]int, 0)
	DFS := func(quees, xyDif, xySum []int) {}
	DFS = func(quees, xyDif, xySum []int) {
		p := len(quees)
		if p == n {
			reslut = append(reslut, quees)
			return
		}
		for q := 0; q < n; q++ {
			if !IsInSlice(quees, q) && !IsInSlice(xyDif, p-q) && !IsInSlice(xySum, p+q) {
				DFS(append(quees, q), append(xyDif, p-q), append(xySum, p+q))
			}
		}
	}

	DFS(make([]int, 0), make([]int, 0), make([]int, 0))
	return bQ(reslut, n)
}

func IsInSlice(res []int, n int) bool {
	for _, value := range res {
		if value == n {
			return true
		}
	}
	return false
}

//
func SolveNQueens2(n int) [][]string {
	if n < 1 {
		return nil
	}

	result := make([][]int, 0)
	xyDiff := make(map[int]interface{}, 0)
	xySum := make(map[int]interface{}, 0)
	cols := make(map[int]interface{}, 0)
	var DFS func(n, row int, curState []int)
	DFS = func(n, row int, curState []int) {
		if row >= n {
			result = append(result, curState)
		}
		for col := 0; col < n; col++ {
			if IsInMap(cols, col) || IsInMap(xySum, row+col) || IsInMap(xyDiff, row-col) {
				continue
			}
			// 增加值
			xyDiff[row-col] = nil
			xySum[row+col] = nil
			cols[col] = nil

			DFS(n, row+1, append(curState, col))

			delete(xySum, row+col)
			delete(xyDiff, row-col)
			delete(cols, col)
		}
	}

	DFS(n, 0, make([]int, 0))
	return bQ(result, n)
}

func IsInMap(res map[int]interface{}, n int) bool {
	_, ok := res[n]
	if ok {
		return true
	}
	return false
}
