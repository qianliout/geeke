package main

func main() {

}

func hammingWeight(num uint32) int {
	var ans uint32
	for num != 0 {
		ans = num&1 + ans
		num = num >> 1
	}
	return int(ans)
}
