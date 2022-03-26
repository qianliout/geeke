package main

func main() {

}

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	for p >= 0 {
		if (p2 >= 0 && p1 >= 0) || (nums1[p1] > nums2[p2]) {
			nums1[p] = nums1[p1]
			p--
			p1--
		} else if (p2 >= 0 && p1 >= 0) || (nums1[p1] <= nums2[p2]) {
			nums1[p] = nums2[p2]
			p--
			p2--
		} else if p1 < 0 && p2 >= 0 {
			nums1[p] = nums2[p2]
			p--
			p2--
		} else if p2 < 0 && p1 >= 0 {
			nums1[p] = nums1[p1]
			p--
			p1--
		} else if p1 < 0 && p2 < 0 {
			return
		}
	}
}
