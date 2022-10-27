package main

func main() {

}

func checkPossibility(nums []int) bool {
	count := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			count++
			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return count <= 1
}
