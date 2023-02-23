package main

func main() {

}

func findPairs(nums []int, k int) int {
	exit := make(map[int]int)
	for i := range nums {
		exit[nums[i]]++
	}

	count := 0
	for _, num := range nums {
		if exit[num] == 0 {
			continue
		}
		if k == 0 {
			if exit[num] > 1 {
				count++
			}
		} else {
			if exit[num-k] > 0 {
				count++
			}
			if exit[num+k] > 0 {
				count++
			}
		}

		exit[num] = 0 // 防止重复计算
	}

	return count
}
