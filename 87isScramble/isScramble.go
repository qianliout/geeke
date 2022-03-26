package main

func main() {

}

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	ss1 := []byte(s1)
	ss2 := []byte(s2)

	n := len(s1)
	dp := make(map[int]map[int]map[int]bool)

	// 初值
	for i := 0; i <= n; i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]map[int]bool)
		}
		for j := 0; j <= n; j++ {
			if dp[i][j] == nil {
				dp[i][j] = make(map[int]bool)
			}
		}
	}

	// 状态方程
	// 这里的l表法起点开始的长度,假设起点是0,长度是1.,那就只有第0个元素,是个右开区间,所以从1开始到n
	for l := 1; l <= n; l++ {
		// S1开始的地方
		for i := 0; i+l <= n; i++ {
			// S2的开始
			for j := 0; j+l <= n; j++ {
				// 如果长度是1,那么是没法分隔的,就只比较原数据就可以了
				if l == 1 {
					dp[l][i][j] = ss1[i] == ss2[j]
				} else {
					// 如果长度不是1,那就要分隔了,
					for p := 1; p < l; p++ {
						dp[l][i][j] = (dp[p][i][j] && dp[l-p][i+p][j+p]) || (dp[p][i][j+l-p] && dp[l-p][i+p][j])
						if dp[l][i][j] == true {
							break
						}
					}
				}
			}
		}

	}
	return dp[n][0][0]
}
