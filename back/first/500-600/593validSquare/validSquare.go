package main

func main() {

}

/*
给定二维空间中四点的坐标，返回四点是否可以构造一个正方形。
一个点的坐标（x，y）由一个有两个整数的整数数组表示。
示例:
输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
输出: True
注意:
所有输入整数都在 [-10000，10000] 范围内。
一个有效的正方形有四个等长的正长和四个等角（90度角）。
输入点没有顺序。
*/
// 一个正方形有两个特点，一，四条边相等，二，任意三个点构成的角是直角
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 四边相等
	s := [][]int{p1, p2, p3, p4}
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			// 这里有个坑，这四个点中，可能会有相同的点
			if s[i][0] == s[j][0] && s[i][1] == s[j][1] {
				return false
			}
			r := (s[i][0]-s[j][0])*(s[i][0]-s[j][0]) + (s[i][1]-s[j][1])*(s[i][1]-s[j][1])
			m[r]++
		}
	}
	if len(m) != 2 {
		return false
	}

	return true
}
