package main

func main() {

}

/*
给定一个二进制数组， 计算其中最大连续1的个数。
示例 1:
输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
*/
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	pre := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == 1 {
			pre++
			if pre > max {
				max = pre
			}
		} else {
			pre = 0
		}
	}
	return max
}
