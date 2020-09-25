package main

func main() {

}

/*
两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。
计算一个数组中，任意两个数之间汉明距离的总和。
示例:
输入: 4, 14, 2
输出: 6
解释: 在二进制表示中，4表示为0100，14表示为1110，2表示为0010。（这样表示是为了体现后四位之间关系）
所以答案为：
HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.
注意:
    数组中元素的范围为从 0到 10^9。
    数组的长度不超过 10^4。
*/

// 会超时
func totalHammingDistance(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			ans += hanming(nums[i], nums[j])
		}
	}
	return ans
}

func hanming(x, y int) int {
	ans := 0
	for x != 0 || y != 0 {
		xlast := x & 1
		ylast := y & 1
		if xlast != ylast {
			ans++
		}
		if x != 0 {
			x >>= 1
		}
		if y != 0 {
			y >>= 1
		}
	}
	return ans
}

// 将所有数转换的成二进制，现在我们先看第0位，求任意两个数第0位的汉明距离的总和。我们从左往右遍历数组，
// 当遍历到第i个数nums[i]时，nums[i]第0位能贡献的汉明距离等于i左边所有数的第0位与nums[i]的第0位不同的数个数，
// 本来还需要加上右边的数，但为了避免重复计算，我们只看左边的数，所以我们可以一边遍历一边统计出现的0和1的个数。
// 这样当我们遍历完数组，就可以得到所有数的第0位能贡献的汉明距离。而题目的答案就是求出所有位能贡献的汉明距离

func totalHamming(nums []int) int {
	ans := 0

	for i := 0; i <= 32; i++ {
		one, zero := 0, 0

		for i := 0; i < len(nums); i++ {
			if nums[i]&1 == 1 {
				ans += zero
				one++
			} else {
				ans += one
				zero++
			}
			nums[i] = nums[i] >> 1
		}
	}
	return ans
}
