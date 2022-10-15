package main

func main() {
	nums := []int{1, 2, 3, 3, 4, 5}
	isPossible(nums)
}

func isPossible(nums []int) bool {
	left := map[int]int{} // 每个数字的剩余个数
	for _, v := range nums {
		left[v]++
	}
	endCnt := map[int]int{} // 以某个数字结尾的连续子序列的个数
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
			left[v]--
			endCnt[v-1]--
			endCnt[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		} else { // 无法生成
			return false
		}
	}
	return true
}
