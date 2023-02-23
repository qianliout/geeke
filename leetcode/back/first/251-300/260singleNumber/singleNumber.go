package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2, 3, 4, 1}
	res := singleNumber(nums)
	fmt.Println("res is ", res)
	fmt.Println(2 ^ -2)
	fmt.Println(3 ^ -3)
}

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
示例 :
输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

	结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
	你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

链接：https://leetcode-cn.com/problems/single-number-iii
*/
func singleNumber(nums []int) []int {
	first := 0
	for _, num := range nums {
		first = first ^ num
	}
	diff := first & (-first) // 得到正数的最后一个1
	x := 0
	for _, num := range nums {
		// 这里等于0还是不等于0都是一样的，主要是将数分成两组，相同的一定是在一个组内，
		// 这样他们异或就是0，最后就只剩下间单独一个的那个数
		if num&diff == 0 {
			x = x ^ num
		}
	}
	return []int{x, first ^ x}
}

func singleNumber1(nums []int) []int {
	first := 0
	for _, n := range nums {
		first = first ^ n
	}
	// 假设两个不同的数是a,b,这样这样first=a^b
	diff := first & (-first) // 得到正数的最后一个1
	// 用diff把nums分成两部分，a和b必定会在两边
	left, right := make([]int, 0), make([]int, 0)
	for _, n := range nums {
		if n&diff == 0 {
			left = append(left, n)
		} else {
			right = append(right, n)
		}
	}
	// 然后就是求值
	a := 0
	for _, n := range left {
		a = a ^ n
	}
	b := 0
	for _, n := range right {
		b = b ^ n
	}
	return []int{a, b}
}
