package main

func main() {

}

/*
给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
设计一个算法使得这 m 个子数组各自和的最大值最小。
示例 1：
输入：nums = [7,2,5,10,8], m = 2
输出：18
解释：
一共有四种方法将 nums 分割为 2 个子数组。 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
*/
func splitArray(nums []int, m int) int {
	sum, max := 0, 0
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if len(nums) == m {
		return max
	}
	// 开始二分
	left, right := max, sum
	for left <= right { // 这里是小于等，下面才会是right = mid -1
		mid := left + (right-left)/2
		s := split(nums, mid)
		// 如果找到了，就再找找，看有没有更小的
		if s == m {
			right = mid - 1
		}
		if s > m {
			left = mid + 1
		}
		if s < m {
			right = mid - 1
		}
	}
	// 这里按正常的代码逻辑应该检查left是否越界，但是这道题目一定不会越界
	return left
}

// 这个函数的意思是：当每个分组的值都不大于maxinter时，最少要分多少个组
func split(nums []int, max int) int {
	ans, sum := 0, 0
	for _, n := range nums {
		if sum+n > max {
			ans++
			sum = n
		} else {
			sum += n
		}
	}
	// 这里的返回值是一个容易出错的点
	return ans + 1
}
