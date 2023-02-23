package main

func main() {

}

/*
给你一个整数数组 nums ，按要求返回一个新数组 counts 。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例 1：

输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/count-of-smaller-numbers-after-self
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	small := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		r := FindSmallIdx(&small, nums[i])
		res[i] = r
	}

	return res
}

// FindSmallIdx 二分法,找到插入左边的位置,并插入左边
func FindSmallIdx(sorted *[]int, target int) int {

	if len(*sorted) == 0 {
		Insert(sorted, target, 0)
		return 0
	}
	start := 0
	end := len(*sorted) - 1

	if (*sorted)[end] < target {
		Insert(sorted, target, end+1)
		return end + 1
	}
	if (*sorted)[start] > target {
		Insert(sorted, target, 0)
		return 0
	}

	for start < end {
		mid := start + (end-start)/2
		if (*sorted)[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	// 插入数据
	Insert(sorted, target, start)
	return start
}

func Insert(a *[]int, c int, i int) []int {
	*a = append((*a)[:i], append([]int{c}, (*a)[i:]...)...)
	return *a
}
