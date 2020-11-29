package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2}
	res := majorityElement(nums)
	fmt.Println("res is ", res)
}

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:
输入: [3,2,3]
输出: 3
示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
链接：https://leetcode-cn.com/problems/majority-element
*/
func majorityElement(nums []int) int {
	// return findUseMap(nums)
	return find3(nums)
}
func findUseMap(nums []int) int {
	exit := make(map[int]int)
	for _, num := range nums {
		exit[num] += 1
		if exit[num] >= len(nums)/2 {
			return num
		}
	}
	return 0
}

// 摩尔投票法思路
func find2(nums []int) int {
	var candidate int
	var count int = 0
	for _, value := range nums {
		if count == 0 {
			candidate = value
		}
		if candidate == value {
			count += 1
		} else {
			count -= 1
		}
	}
	// 摩尔投票法还有一步是进行检测，这个数是否超过一半，因为题目中说了一定存在，所以就少了这一步检查
	return candidate
}

/*
由于众数数组中出现次数大于n/2 ，那么众数对应的二进制每一个为1的位数出现的次数一定大于n/2，由此可以得出众数
这个方法可以得到正确的结果，但是不容易想到
*/
func find3(nums []int) int {
	result := 0
	k := len(nums) >> 1
	// 注意，这里如果是32位的话，就会溢出，所以选择64位
	for j := 0; j < 64; j++ {
		count := 0

		for _, num := range nums {
			count += num >> j & 1
			if count > k {
				result += 1 << j
				break
			}
		}
	}
	return result
}
