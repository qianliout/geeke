package main

func main() {

}

func containsDuplicate(nums []int) bool {
	exit := make(map[int]int)
	for i := range nums {
		if exit[nums[i]] > 0 {
			return true
		}
		exit[nums[i]]++
	}
	return false
}
