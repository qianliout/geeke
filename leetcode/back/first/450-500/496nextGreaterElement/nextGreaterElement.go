package main

func main() {

}

/*
给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
示例 1:
输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
示例 2:
输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
    对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
提示：
nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。
*/

// 单调栈的方法,这里一定要理解题意，不是说index的右边，而是值的右边
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stark := make([]int, 0)
	hash := make(map[int]int)

	ans := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stark) > 0 && stark[len(stark)-1] <= nums2[i] {
			stark = stark[:len(stark)-1] // 把比num[i]小的数干掉
		}
		// 保存下想要的结果
		if len(stark) == 0 {
			hash[nums2[i]] = -1
		} else {
			hash[nums2[i]] = stark[len(stark)-1]
		}
		// 把num[i]加到stark中，接受后面的判断
		stark = append(stark, nums2[i])
	}
	for _, n := range nums1 {
		ans = append(ans, hash[n])
	}
	return ans
}
