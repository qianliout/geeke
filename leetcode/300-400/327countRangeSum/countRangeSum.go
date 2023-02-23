package main

func main() {

}

// O(n2) 会超时
func countRangeSum(nums []int, lower int, upper int) int {
	res := 0
	if len(nums) == 0 || lower > upper {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		ans := nums[i]
		for j := i; j < len(nums); j++ {
			if j != i {
				ans += nums[j]
			}
			if ans >= lower && ans <= upper {
				res++
			}
		}
	}
	return res
}
