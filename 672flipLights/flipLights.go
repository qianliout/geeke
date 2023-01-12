package main

func main() {

}

func flipLights(n, presses int) int {
	seen := map[int]struct{}{}
	for i := 0; i < 1<<4; i++ {
		pressArr := [4]int{}
		sum := 0
		for j := 0; j < 4; j++ {
			pressArr[j] = i >> j & 1
			sum += pressArr[j]
		}
		if sum%2 == presses%2 && sum <= presses {
			status := pressArr[0] ^ pressArr[2] ^ pressArr[3]
			if n >= 2 {
				status |= (pressArr[0] ^ pressArr[1]) << 1
			}
			if n >= 3 {
				status |= (pressArr[0] ^ pressArr[2]) << 2
			}
			if n >= 4 {
				status |= (pressArr[0] ^ pressArr[1] ^ pressArr[3]) << 3
			}
			seen[status] = struct{}{}
		}
	}
	return len(seen)
}
