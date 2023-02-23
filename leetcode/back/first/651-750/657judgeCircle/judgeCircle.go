package main

func main() {

}

func judgeCircle(moves string) bool {
	move := make(map[byte][]int)
	move['R'] = []int{1, 0}
	move['L'] = []int{-1, 0}
	move['U'] = []int{0, 1}
	move['D'] = []int{0, -1}

	x, y := 0, 0
	ss := []byte(moves)
	for i := 0; i < len(moves); i++ {
		b := ss[i]
		x += move[b][0]
		y += move[b][1]
	}
	return x == 0 && y == 0
}
