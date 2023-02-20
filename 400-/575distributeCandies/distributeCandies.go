package main

func main() {

}

func distributeCandies(candyType []int) int {
	exit := make(map[int]bool)
	for i := range candyType {
		exit[candyType[i]] = true
	}
	ans := len(exit)
	if len(candyType)/2 < ans {
		ans = len(candyType) / 2
	}
	return ans
}
